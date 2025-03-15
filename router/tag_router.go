package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) TagRouter() {
	tag := api.ApiGroupApp.TagApi
	router.POST("tags", middleware.JwtAdmin(), tag.TagCreateView)

	router.GET("tags", tag.TagListView)

	router.GET("tag_names", tag.TagNameListView)

	router.PUT("tags/:id", middleware.JwtAdmin(), tag.TagUpdateView)

	router.DELETE("tags", middleware.JwtAdmin(), tag.TagRemoveView)

}
