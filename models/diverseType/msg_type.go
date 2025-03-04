package diverseType

import "encoding/json"

type MsgType int

const (
	InRoomMsg  MsgType = 1 // 进入聊天室
	TextMsg    MsgType = 2 // 文本消息
	ImageMsg   MsgType = 3 // 图片消息
	VoiceMsg   MsgType = 4 // 语音消息
	VideoMsg   MsgType = 5 // 视频消息
	SystemMsg  MsgType = 6 // 系统消息
	OutRoomMsg MsgType = 7 // 退出聊天室
)

func (m MsgType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m MsgType) String() string {
	var str string
	switch m {
	case InRoomMsg:
		str = "进入聊天室"
	case TextMsg:
		str = "文本消息"
	case ImageMsg:
		str = "图片消息"
	case VoiceMsg:
		str = "语音消息"
	case VideoMsg:
		str = "语音消息"
	case SystemMsg:
		str = "系统消息"
	case OutRoomMsg:
		str = "退出聊天室"
	}
	return str
}
