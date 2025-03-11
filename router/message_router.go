package router

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) MessageRouter() {
	msg := api.ApiGroupApp.MessageApi
	router.POST("messages", middleware.JwtAuth(), msg.MessageCreateView)

	// =================== 聊天信息(不含记录) =================
	// 某人收到不同发送者消息的列表
	router.GET("message_users", middleware.JwtAdmin(), msg.MessageUserListView)
	// 某个用户的聊天列表，包含与该用户有聊天记录的其他用户信息以及他们之间的消息数量
	router.GET("message_users/user", middleware.JwtAdmin(), msg.MessageUserListByUserView)
	// 我的消息列表
	router.GET("message_users/me", middleware.JwtAuth(), msg.MessageUserListByMeView)

	// =================== 聊天记录 =================
	// 两个用户之间的聊天列表
	router.GET("message_users/record", middleware.JwtAdmin(), msg.MessageUserRecordView)
	// 我与某个用户的聊天列表
	router.GET("message_users/record/me", middleware.JwtAuth(), msg.MessageUserRecordByMeView)
	// 所有消息列表
	router.GET("messages_all", msg.MessageListAllView)
	// 与自己相关的消息列表
	router.GET("messages", middleware.JwtAuth(), msg.MessageListView)

	router.DELETE("messages", middleware.JwtAuth(), msg.MessageRemoveView)
}
