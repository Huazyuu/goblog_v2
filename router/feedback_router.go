package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) FeedbackApi() {
	fb := api.ApiGroupApp.FeedbackApi
	router.POST("feedback", fb.FeedBackCreateView)
	router.GET("feedback", fb.FeedBackListView)
	router.DELETE("feedback", middleware.JwtAdmin(), fb.FeedBackRemoveView)
}
