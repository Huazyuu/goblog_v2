package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) AdvertRouter() {
	advertsApi := api.ApiGroupApp.AdvertsApi
	router.POST("adverts", middleware.JwtAdmin(), advertsApi.AdvertCreateView)
	router.GET("adverts", middleware.JwtAdmin(), advertsApi.AdvertListView)
	router.PUT("adverts/:id", middleware.JwtAdmin(), advertsApi.AdvertUpdateView)
	router.DELETE("adverts", middleware.JwtAdmin(), advertsApi.AdvertRemoveView)

}
