package chat_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

func CreateChatMsg(chat *sqlmodels.ChatModel) {
	global.DB.Create(chat)
}
