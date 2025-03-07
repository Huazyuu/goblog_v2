package api

import (
	"backend/api/adverts_api"
	"backend/api/images_api"
	"backend/api/settings_api"
	"backend/api/users_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	UsersApi    users_api.UsersApi
	AdvertsApi  adverts_api.AdvertsApi
}

var ApiGroupApp = new(ApiGroup)
