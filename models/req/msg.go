package req

import "backend/models/common"

type MessageRequest struct {
	RevUserID uint   `json:"rev_user_id" binding:"required"` // 接收人id
	Content   string `json:"content" binding:"required"`     // 消息内容
}
type MessageUserListRequest struct {
	common.PageInfo
	NickName string `json:"nickName" form:"nickName"`
}

type MessageUserListByUserRequest struct {
	UserID uint `json:"userID" form:"userID" binding:"required"`
}

type MessageUserRecordRequest struct {
	common.PageInfo
	SendUserID uint `json:"sendUserID" form:"sendUserID" binding:"required"`
	RevUserID  uint `json:"revUserID" form:"revUserID" binding:"required"`
}
type MessageUserRecordByMeRequest struct {
	common.PageInfo
	UserID uint `json:"userID" form:"userID" binding:"required"`
}
