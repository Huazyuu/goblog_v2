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
