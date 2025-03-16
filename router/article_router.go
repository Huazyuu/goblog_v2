package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) ArticleRouter() {
	article := api.ApiGroupApp.ArticleApi
	// 发布文章
	router.POST("articles", middleware.JwtAuth(), article.ArticleCreateView)

	// ES article look++
	{
		// 文章内容
		router.GET("articles/detail/:id", article.ArticleDetailByIDView)
		// 通过title(精准查找)找出具体文章
		router.GET("article/detail", article.ArticleDetailByTitleView)
		// 文章详情
		router.GET("articles/:id", article.ArticleDetailView)
	}

	// 文章列表 全
	router.GET("articles", article.ArticleListView)
	// 文章 id - title 列表
	router.GET("article_id_title", article.ArticleIDTitleListView)
	// 文章 category - count 列表
	router.GET("article/categories", article.ArticleCategoryListView)
	// 文章标签列表 tag count [title]
	router.GET("articles/tags", article.ArticleTagListView)
	// 文章发布日历 data count
	router.GET("articles/calendar", article.ArticleCalendarView)

	// 更新文章
	router.PUT("articles", middleware.JwtAuth(), article.ArticleUpdateView)
	// 删除文章
	router.DELETE("articles", middleware.JwtAuth(), article.ArticleRemoveView)

	// 文章收藏
	router.POST("articles/collects/:id", middleware.JwtAuth(), article.ArticleCollCreateView)
	// 用户收藏的文章列表
	router.GET("articles/collects", middleware.JwtAuth(), article.ArticleCollListView)
	// 批量撤销收藏
	router.DELETE("articles/collects", middleware.JwtAuth(), article.ArticleCollRemoveView)

	// 文章点赞
	router.POST("articles/digg", middleware.JwtAuth(), article.ArticleDiggView)

}
