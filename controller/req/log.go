package req

import (
	"backend/models/diverseType"
)

type LogListRequest struct {
	PageInfo
	Level    diverseType.LogLevel `json:"level" form:"level"`       // 日志查询的等级
	Type     diverseType.LogType  `json:"type" form:"type"`         // 日志的类型   1 登录日志  2 操作日志  3 运行日志
	IP       string               `json:"ip" form:"ip"`             // 根据ip查询
	UserID   uint                 `json:"userID" form:"user_id"`    // 根据用户id查询
	Addr     string               `json:"addr" form:"addr"`         // 感觉地址查询
	Date     string               `json:"date" form:"date"`         // 查某一天的，格式是年月日
	Status   *bool                `json:"status" form:"status"`     // 登录状态查询  true  成功  false 失败
	UserName string               `json:"userName" form:"username"` // 查用户名
}
