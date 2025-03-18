package req

import (
	"backend/models/diverseType"
	"time"
)

type GroupRequest struct {
	Content string              `json:"content"`  // 聊天的内容
	MsgType diverseType.MsgType `json:"msg_type"` // 聊天类型
}
type GroupResponse struct {
	NickName    string              `json:"nick_name"`    // 前端自己生成
	Avatar      string              `json:"avatar"`       // 头像
	MsgType     diverseType.MsgType `json:"msg_type"`     // 聊天类型
	Content     string              `json:"content"`      // 聊天的内容
	OnlineCount int                 `json:"online_count"` // 在线人数
	Date        time.Time           `json:"created_at"`   // 消息的时间
}
