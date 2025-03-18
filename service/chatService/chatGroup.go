package chatService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"backend/repository/chat_repo"
	"backend/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strings"
	"sync"
	"time"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var connGroupMap = map[string]ChatUser{}
var mutex sync.Mutex

func HandleConnection(conn *websocket.Conn) {
	defer handleDisConnect(conn)

	addr := conn.RemoteAddr().String()
	user := createChatUser(conn)

	mutex.Lock()
	connGroupMap[addr] = user
	mutex.Unlock()

	global.Log.Infof("%s connected as %s", addr, user.NickName)

	notifyUserEntry(user)
	handleMessages(user)
}

func handleMessages(user ChatUser) {
	for {
		_, msg, err := user.Conn.ReadMessage()
		if err != nil {
			break
		}
		var request req.GroupRequest
		if err = json.Unmarshal(msg, &request); err != nil {
			sendSystemMessage(user, "消息解析失败", true)
			continue
		}
		processMessage(user, request)
	}
}

func handleDisConnect(conn *websocket.Conn) {
	addr := conn.RemoteAddr().String()
	mutex.Lock()
	user, exist := connGroupMap[addr]
	delete(connGroupMap, addr)
	mutex.Unlock()
	if exist {
		notifyUserExit(user)
	}
	conn.Close()
}

func processMessage(user ChatUser, request req.GroupRequest) {
	switch request.MsgType {
	case diverseType.TextMsg:
		handleTextMessage(user, request.Content)
	default:
		sendSystemMessage(user, "不支持的消息类型", true)
	}
}

func createChatUser(conn *websocket.Conn) ChatUser {
	nickName := utils.GenerateName()
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", string([]rune(nickName)[0]))
	return ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
}

func notifyUserEntry(user ChatUser) {
	sendSystemMessage(user, "进入聊天室", false)
	broadcastMessage(req.GroupResponse{
		NickName:    user.NickName,
		Avatar:      user.Avatar,
		MsgType:     diverseType.InRoomMsg,
		Content:     fmt.Sprintf("%s 进入聊天室", user.NickName),
		OnlineCount: len(connGroupMap),
		Date:        time.Now(),
	}, user, true)
}

func notifyUserExit(user ChatUser) {
	broadcastMessage(req.GroupResponse{
		NickName:    user.NickName,
		Avatar:      user.Avatar,
		MsgType:     diverseType.OutRoomMsg,
		Content:     fmt.Sprintf("%s 退出聊天室", user.NickName),
		OnlineCount: len(connGroupMap),
		Date:        time.Now(),
	}, user, true)
}

func broadcastMessage(response req.GroupResponse, user ChatUser, isGroup bool) {
	msgBytes, _ := json.Marshal(response)
	mutex.Lock()
	defer mutex.Unlock()
	for _, user := range connGroupMap {
		if err := user.Conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			global.Log.Error("发送消息失败: %v", err)
		}
	}
	if isGroup {
		saveGroupMessage(user, response)
	}
}

func saveGroupMessage(user ChatUser, response req.GroupResponse) {
	ip, addr := getIPAndPort(user.Conn.RemoteAddr().String())
	chat_repo.CreateChatMsg(&sqlmodels.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
		MsgType:  response.MsgType,
	})
}

func sendSystemMessage(user ChatUser, content string, save bool) {
	response := req.GroupResponse{
		NickName:    user.NickName,
		Avatar:      user.Avatar,
		MsgType:     diverseType.SystemMsg,
		Content:     content,
		OnlineCount: len(connGroupMap),
		Date:        time.Now(),
	}
	msgBytes, _ := json.Marshal(response)
	user.Conn.WriteMessage(websocket.TextMessage, msgBytes)
	if save {
		saveSystemMessage(user, content)
	}
}

func saveSystemMessage(user ChatUser, content string) {
	ip, addr := getIPAndPort(user.Conn.RemoteAddr().String())
	chat_repo.CreateChatMsg(&sqlmodels.ChatModel{
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Content:  content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
		MsgType:  diverseType.SystemMsg,
	})
}

func handleTextMessage(user ChatUser, content string) {
	if strings.TrimSpace(content) == "" {
		sendSystemMessage(user, "消息不能为空", false)
		return
	}
	broadcastMessage(req.GroupResponse{
		NickName:    user.NickName,
		Avatar:      user.Avatar,
		MsgType:     diverseType.TextMsg,
		Content:     content,
		OnlineCount: len(connGroupMap),
		Date:        time.Now(),
	}, user, true)
}
func getIPAndPort(addr string) (ip string, port string) {
	addrList := strings.Split(addr, ":")
	ip = addrList[0]
	port = utils.GetAddr(ip)
	return ip, port
}
