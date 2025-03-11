package msgService

import (
	"backend/middleware/jwt"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"backend/repository/msg_repo"
	"backend/utils"
)

func MsgList(claims *jwt.CustomClaims) ([]diverseType.Message, error) {
	var messageGroup = make(diverseType.MessageGroup) // 用于存储分组后的消息
	var messageList []sqlmodels.MessageModel          // 用于存储从数据库中查询到的所有消息记录
	var messages = make([]diverseType.Message, 0)     // 用于存储最终要返回的消息列表

	messageList, err := msg_repo.GetMsgListRela(claims.UserID, claims.UserID)
	if err != nil {
		return messages, err
	}
	for _, msgModel := range messageList {
		msg := diverseType.Message{
			SendUserID:       msgModel.SendUserID,
			SendUserNickName: msgModel.SendUserNickName,
			SendUserAvatar:   msgModel.SendUserAvatar,
			RevUserID:        msgModel.RevUserID,
			RevUserNickName:  msgModel.RevUserNickName,
			RevUserAvatar:    msgModel.RevUserAvatar,
			Content:          msgModel.Content,
			CreatedAt:        msgModel.CreatedAt,
			MessageCount:     1,
		}
		idStr := utils.CombineIDs(msgModel.SendUserID, msgModel.RevUserID)
		val, ok := messageGroup[idStr]
		if !ok {
			// 不存在
			messageGroup[idStr] = &msg
			continue
		}
		msg.MessageCount = val.MessageCount + 1
		messageGroup[idStr] = &msg
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}
	return messages, nil
}
