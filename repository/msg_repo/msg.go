package msg_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

type MessageUserStat struct {
	SendUserID uint
	RevUserID  uint
	Count      int
}

func CreateMessage(msg sqlmodels.MessageModel) error {
	return global.DB.Create(&msg).Error
}

func GetMsgListRela(senderID, receiverID uint) (messageList []sqlmodels.MessageModel, err error) {
	result := global.DB.Order("created_at asc").Where("send_user_id = ? or rev_user_id = ?", senderID, receiverID).Find(&messageList)
	return messageList, result.Error
}

func CountMessageUsers(nickname string) (count int64, err error) {
	query := global.DB.Model(&sqlmodels.MessageModel{})
	if nickname != "" {
		query = query.Where("send_user_nick_name = ?", nickname)
	}
	err = query.Group("send_user_id").Count(&count).Error
	return count, err
}

// GetMessageUserStats 根据用户id聚合,查看对应昵称的用户给多少人发送了消息
func GetMessageUserStats(nickname string, page, limit int) (stats []MessageUserStat, err error) {
	offset := (page - 1) * limit
	query := global.DB.Model(&sqlmodels.MessageModel{})
	if nickname != "" {
		query = query.Where("send_user_nick_name = ?", nickname)
	}

	err = query.Select("send_user_id, count(distinct rev_user_id) as count").
		Group("send_user_id").
		Offset(offset).
		Limit(limit).
		Scan(&stats).Error
	/*
			SELECT send_user_id, count(distinct rev_user_id) as count
		 		FROM `message`
				WHERE send_user_nick_name = '测试用户'
				GROUP BY `send_user_id`
				LIMIT 10
	*/
	return stats, err
}

func GetUserMessageStats(userID uint) (stats []MessageUserStat, err error) {
	err = global.DB.Model(&sqlmodels.MessageModel{}).
		Where("send_user_id = ? OR rev_user_id = ?", userID, userID).
		Select("send_user_id", "rev_user_id", "count(id) as count").
		Group("send_user_id, rev_user_id").
		Scan(&stats).Error
	return stats, err
}

func GetMsgsByIDList(idList []uint) (msgLists []sqlmodels.MessageModel, err error) {
	err = global.DB.Find(&msgLists, idList).Error
	return msgLists, err
}
func GetMyMsgsByIDList(userID uint, idList []uint) (msgLists []sqlmodels.MessageModel, err error) {
	err = global.DB.Where("id IN (?) AND send_user_id = ?", idList, userID).
		Find(&msgLists).Error
	return msgLists, err
}
func DeleteMsgs(msgs []sqlmodels.MessageModel) (count int64, err error) {
	res := global.DB.Delete(&msgs)
	count = res.RowsAffected
	err = res.Error
	return count, err
}
