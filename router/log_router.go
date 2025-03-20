package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) LogRouter() {
	logRouter := api.ApiGroupApp.LogApi
	router.GET("logs", middleware.JwtAdmin(), logRouter.LogListView)

	router.GET("logs_refresh", middleware.JwtAdmin(), logRouter.LogRefreshView)

	router.DELETE("logs", middleware.JwtAdmin(), logRouter.LogRemoveListView)
}
