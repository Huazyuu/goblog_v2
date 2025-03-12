package router

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) ArticleRouter() {
	article := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAdmin(), article.ArticleCreateView)
}
