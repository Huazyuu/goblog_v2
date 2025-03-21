package data_repo

import (
	"backend/global"
	"fmt"
	"time"
)

func GetCountByDate(tableName string, startTime, endTime time.Time) (map[string]int, error) {
	var result []struct {
		Date  string
		Count int
	}

	err := global.DB.Table(tableName).
		Select("DATE_FORMAT(created_at, '%Y-%m-%d') as date", "COUNT(id) as count").
		Where("created_at BETWEEN ? AND ?", startTime, endTime).
		Group("date").
		Scan(&result).Error

	if err != nil {
		return nil, fmt.Errorf("数据库查询失败: %v", err)
	}

	countMap := make(map[string]int)
	for _, v := range result {
		countMap[v.Date] = v.Count
	}
	return countMap, nil
}

// type DataCount struct {
// 	UserCount      int `json:"user_count"`
// 	ArticleCount   int `json:"article_count"`
// 	MessageCount   int `json:"message_count"`
// 	ChatGroupCount int `json:"chat_group_count"`
// 	NowLoginCount  int `json:"now_login_count"`
// 	NowSignCount   int `json:"now_sign_count"`
// 	FlowCount      int `json:"flow_count"`
// }
//
// func GetDataCount() (cnts DataCount, err error) {
// 	result, err := global.ESClient.
// 		Search(esmodels.ArticleModel{}.Index()).
// 		Query(elastic.NewMatchAllQuery()).
// 		Do(context.Background())
// 	if err != nil {
// 		return cnts, err
// 	}
// 	cnts.ArticleCount = int(result.Hits.TotalHits.Value) // 搜索到结果总条数
// 	err = global.DB.Model(sqlmodels.UserModel{}).Select("count(id)").Scan(&cnts.UserCount).Error
// 	if err != nil {
// 		return cnts, err
// 	}
// 	err = global.DB.Model(sqlmodels.MessageModel{}).Select("count(id)").Scan(&cnts.MessageCount).Error
// 	if err != nil {
// 		return cnts, err
// 	}
// 	err = global.DB.Model(sqlmodels.ChatModel{IsGroup: true}).Select("count(id)").Scan(&cnts.ChatGroupCount).Error
// 	if err != nil {
// 		return cnts, err
// 	}
// 	err = global.DB.Model(sqlmodels.LoginDataModel{}).Where("to_days(created_at)=to_days(now())").
// 		Select("count(id)").Scan(&cnts.NowLoginCount).Error
// 	if err != nil {
// 		return cnts, err
// 	}
// 	err = global.DB.Model(sqlmodels.UserModel{}).Where("to_days(created_at)=to_days(now())").
// 		Select("count(id)").Scan(&cnts.NowSignCount).Error
// 	if err != nil {
// 		return cnts, err
// 	}
// 	return cnts, nil
// }
