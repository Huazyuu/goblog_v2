package router

import (
	"backend/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	// 查看网站信息
	router.GET("settings/site", settingsApi.SettingsSiteInfoView)
	// 修改网站信息
	// todo 中间件添加 middleware.auth
	router.PUT("settings/site", settingsApi.SettingsSiteUpdateView)
	// 获取某一项信息
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	// 修改某一项信息
	// todo 中间件添加 middleware.auth
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
