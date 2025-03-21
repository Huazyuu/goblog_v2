package dataService

import (
	"backend/global"
	"backend/repository/data_repo"
)

type DataSumResponse struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	ChatGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
	FlowCount      int `json:"flow_count"`
}

func DataSumService() (counts data_repo.DataCount, err error) {
	counts, err = data_repo.NewDataService(global.DB, global.ESClient).GetCounts()
	if err != nil {
		return
	}
	return
}
