package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) ArticleRouter() {
	article := api.ApiGroupApp.ArticleApi
	// 发布文章
	router.POST("articles", middleware.JwtAdmin(), article.ArticleCreateView)
	// 文章列表
	router.GET("articles", article.ArticleListView)
	// 文章 id - title 列表
	router.GET("article_id_title", article.ArticleIDTitleListView)
	// 文章 category - count 列表
	router.GET("article/categories", article.ArticleCategoryListView)
	// 通过title(精准查找)找出具体文章
	router.GET("article/detail", article.ArticleDetailByTitleView)
	// 文章发布日历
	router.GET("articles/calendar", article.ArticleCalendarView)
	// 文章标签列表
	router.GET("articles/tags", article.ArticleTagListView)
}
