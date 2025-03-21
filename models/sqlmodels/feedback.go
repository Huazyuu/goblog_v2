package sqlmodels

// FeedbackModel 反馈
type FeedbackModel struct {
	MODEL
	Email   string `gorm:"size:64" json:"email"`
	Content string `gorm:"size:256" json:"content"`
}

func (FeedbackModel) TableName() string {
	return "feedback"
}
