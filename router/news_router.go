package router

import (
	"backend/controller/api"
)

func (router RouterGroup) NewsRouter() {
	news := api.ApiGroupApp.NewsApi
	router.GET("news", news.NewListView)
}
