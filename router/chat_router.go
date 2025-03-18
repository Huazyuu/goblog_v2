package router

import "backend/controller/api"

func (router RouterGroup) ChatRouter() {
	chat := api.ApiGroupApp.ChatApi
	router.GET("chat_groups", chat.ChatGroupView)

}
