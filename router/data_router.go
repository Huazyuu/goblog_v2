package router

import (
	"backend/controller/api"
)

func (router RouterGroup) DataRouter() {
	data := api.ApiGroupApp.DataApi
	router.GET("data_login", data.LoginDataView)

	router.GET("data_sum", data.DataSumView)

	router.GET("role_ids", data.RoleIDListView)

}
