package req

type MessageRequest struct {
	RevUserID uint   `json:"rev_user_id" binding:"required"` // 接收人id
	Content   string `json:"content" binding:"required"`     // 消息内容
}
type MessageUserListRequest struct {
	PageInfo
	NickName string `json:"nickName" form:"nickName"`
}

type MessageUserListByUserRequest struct {
	UserID uint `json:"userID" form:"userID" binding:"required"`
}

type MessageUserRecordRequest struct {
	PageInfo
	SendUserID uint `json:"sendUserID" form:"sendUserID" binding:"required"`
	RevUserID  uint `json:"revUserID" form:"revUserID" binding:"required"`
}
type MessageUserRecordByMeRequest struct {
	PageInfo
	UserID uint `json:"userID" form:"userID" binding:"required"`
}
