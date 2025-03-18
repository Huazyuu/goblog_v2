package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) ChatRouter() {
	chat := api.ApiGroupApp.ChatApi
	router.GET("chat/groups", chat.ChatGroupView)
	router.GET("chat/records", chat.ChatListView)
	router.DELETE("chat/records", middleware.JwtAdmin(), chat.ChatRemoveView)

}
