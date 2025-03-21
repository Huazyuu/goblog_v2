package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) GaodeRouter() {
	gaode := api.ApiGroupApp.GaodeApi
	router.GET("gaode/ip_weather", middleware.JwtAuth(), gaode.WeatherInfoByIPView)
	router.GET("gaode/adcode_weather", middleware.JwtAuth(), gaode.WeatherInfoByAdcodeView)

}
