package logService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/sqlmodels"
	"time"
)

func LogListService(cr req.LogListRequest) ([]sqlmodels.LogModel, string, error) {
	var query = global.DB.Where("")
	if cr.Date != "" {
		_, dateTimeErr := time.Parse("2006-01-02", cr.Date)
		if dateTimeErr != nil {
			return nil, "时间格式错误", dateTimeErr
		}
		query.Where("date(created_at) = ?", cr.Date)
	}
	if cr.Status != nil {
		query.Where("status = ?", cr.Status)
	}

	list, _, err := req.ComList(sqlmodels.LogModel{
		Type:     cr.Type,
		Level:    cr.Level,
		IP:       cr.IP,
		Addr:     cr.Addr,
		UserID:   cr.UserID,
		UserName: cr.UserName,
	}, req.Option{
		PageInfo: cr.PageInfo,
		Where:    query,
		Likes:    []string{"title", "user_name"},
	})
	if err != nil {
		return nil, "查询错误", err
	}
	return list, "查询成功", nil

}
