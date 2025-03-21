package req

type FeedBackCreateRequest struct {
	Email   string `json:"email" binding:"required,email" msg:"请输入邮箱"`
	Content string `json:"content" binding:"required" msg:"请输入内容"`
}
