package resp

type DataSumResponse struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	ChatGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
	FlowCount      int `json:"flow_count"`
}

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}
