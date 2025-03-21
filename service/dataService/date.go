package dataService

import (
	"backend/controller/req"
	"backend/controller/resp"
	"backend/models/sqlmodels"
	"backend/repository/data_repo"
	"time"
)

var dateDurationMap = map[req.DateType]time.Duration{
	req.OneWeek:    7 * 24 * time.Hour,
	req.OneMonth:   30 * 24 * time.Hour,
	req.TwoMonth:   60 * 24 * time.Hour,
	req.ThreeMonth: 90 * 24 * time.Hour,
	req.HalfYear:   180 * 24 * time.Hour,
	req.OneYear:    365 * 24 * time.Hour,
}

func LoginDataService(cr req.DateRequest) (res resp.DateCountResponse, err error) {
	// 获取时间范围
	endTime := time.Now()
	var startTime time.Time
	if cr.Date != 0 {
		startTime = endTime.Add(-dateDurationMap[cr.Date])
	} else {
		startTime = endTime.Add(-dateDurationMap[req.OneWeek])
	}

	// 获取登录数据
	loginData, err := data_repo.GetCountByDate(sqlmodels.LoginDataModel{}.TableName(), startTime, endTime)
	if err != nil {
		return res, err
	}

	// 获取注册数据
	signData, err := data_repo.GetCountByDate(sqlmodels.UserModel{}.TableName(), startTime, endTime)
	if err != nil {
		return res, err
	}

	// 填充数据
	res.DateList = generateDateList(startTime, endTime)
	res.LoginData = fillCountData(res.DateList, loginData)
	res.SignData = fillCountData(res.DateList, signData)

	return res, nil
}
func generateDateList(start, end time.Time) []string {
	var dates []string
	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}
	return dates
}

func fillCountData(dateList []string, data map[string]int) []int {
	var counts []int
	for _, date := range dateList {
		counts = append(counts, data[date])
	}
	return counts
}
