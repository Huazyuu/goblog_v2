package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) CommentRouter() {
	comment := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), comment.CommentCreateView)
	router.GET("comments", comment.CommentListView)
	// 有评论的文章列表
	router.GET("comments/articles", middleware.JwtAdmin(), comment.CommentByArticleListView)
	// 文章下的评论列表
	router.GET("comments/:id", comment.CommentListInArticleView)
	// 评论点赞
	router.GET("comments/digg/:id", comment.CommentDiggView)
	// 删除评论
	router.DELETE("comments/:id", middleware.JwtAuth(), comment.CommentRemoveView)
}
