package router

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	// 查看网站信息
	router.GET("settings/site", settingsApi.SettingsSiteInfoView)
	// 修改网站信息
	router.PUT("settings/site", middleware.JwtAdmin(), settingsApi.SettingsSiteUpdateView)
	// 获取某一项信息
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	// 修改某一项信息
	router.PUT("settings/:name", middleware.JwtAdmin(), settingsApi.SettingsInfoUpdateView)
}
