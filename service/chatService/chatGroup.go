package chatService

import (
	"backend/controller/req"
	"backend/global"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
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
			// 无法读取消息,执行退出
			global.Log.Errorf("Read message error: %v", err)
			break
		}
		var request req.GroupRequest
		if err = json.Unmarshal(msg, &request); err != nil {
			sendSystemMessage(user, "消息解析失败", false)
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
