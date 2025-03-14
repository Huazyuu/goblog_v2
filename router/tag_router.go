package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) TagRouter() {
	tag := api.ApiGroupApp.TagApi
	router.POST("tags", middleware.JwtAdmin(), tag.TagCreateView)

	router.GET("tags", tag.TagListView)
	// todo tag list_name router 需要查询es 找出对应tag
	router.GET("tag_names", tag.TagNameListView)

	router.PUT("tags/:id", middleware.JwtAdmin(), tag.TagUpdateView)

	router.DELETE("tags", middleware.JwtAdmin(), tag.TagRemoveView)

}
