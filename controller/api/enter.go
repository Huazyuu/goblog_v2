package api

import (
	"backend/controller/api/adverts_api"
	"backend/controller/api/article_api"
	"backend/controller/api/chat_api"
	"backend/controller/api/comment_api"
	"backend/controller/api/data_api"
	"backend/controller/api/gaode_api"
	"backend/controller/api/images_api"
	"backend/controller/api/log_api"
	"backend/controller/api/menu_api"
	"backend/controller/api/message_api"
	"backend/controller/api/news_api"
	"backend/controller/api/settings_api"
	"backend/controller/api/tag_api"
	"backend/controller/api/users_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	UsersApi    users_api.UsersApi
	AdvertsApi  adverts_api.AdvertsApi
	MenuApi     menu_api.MenuApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	CommentApi  comment_api.CommentApi
	NewsApi     news_api.NewsApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	DataApi     data_api.DataApi
	GaodeApi    gaode_api.GaodeApi
}

var ApiGroupApp = new(ApiGroup)
