package chat_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"errors"
)

func CreateChatMsg(chat *sqlmodels.ChatModel) {
	global.DB.Create(chat)
}
func RemoveChatMsg(ids []uint) (error, int) {
	var chatList []sqlmodels.ChatModel
	global.DB.Find(&chatList, ids)

	if len(chatList) > 0 {
		err := global.DB.Delete(&chatList).Error
		if err != nil {
			return err, 0
		}
	} else {
		return errors.New("没有这条消息"), 0
	}
	return nil, len(chatList)
}
