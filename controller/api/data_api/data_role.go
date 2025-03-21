package data_api

import (
	"backend/controller/res"
	"github.com/gin-gonic/gin"
)

func (DataApi) RoleIDListView(c *gin.Context) {
	type T struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	}
	res.OkWithData([]T{
		{"管理员", 1},
		{"普通用户", 2},
		{"游客", 3},
	}, c)
}
