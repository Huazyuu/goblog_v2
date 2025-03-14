package api

import (
	"backend/controller/api/adverts_api"
	"backend/controller/api/article_api"
	"backend/controller/api/images_api"
	"backend/controller/api/menu_api"
	"backend/controller/api/message_api"
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
}

var ApiGroupApp = new(ApiGroup)
