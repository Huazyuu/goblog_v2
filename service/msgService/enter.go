package msgService

type MessageUserListResponse struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	UserID   uint   `json:"userID"`
	Avatar   string `json:"avatar"`
	Count    int    `json:"count"`
}
