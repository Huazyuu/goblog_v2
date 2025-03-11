package msgService

import (
	"backend/models/req"
	"backend/models/sqlmodels"
	"backend/repository/msg_repo"
	"backend/repository/user_repo"
)

func MsgCreateService(cr req.MessageRequest, userid uint) (string, error) {
	var sender, receiver sqlmodels.UserModel
	sender, err := user_repo.GetByID(userid)
	if err != nil {
		return "发送者不存在", err
	}
	receiver, err = user_repo.GetByID(cr.RevUserID)
	if err != nil {
		return "接收者不存在", err
	}
	err = msg_repo.CreateMessage(sqlmodels.MessageModel{
		SendUserID:       sender.ID,
		SendUserNickName: sender.NickName,
		SendUserAvatar:   sender.Avatar,
		RevUserID:        receiver.ID,
		RevUserNickName:  receiver.NickName,
		RevUserAvatar:    receiver.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	})
	if err != nil {
		return "发送消息失败", err
	}
	return "发送消息成功", nil

}
