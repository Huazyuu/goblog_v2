package router

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("image", middleware.JwtAuth(), imagesApi.ImageUploadView)
	router.POST("images", middleware.JwtAuth(), imagesApi.ImagesUploadView)
	router.POST("freeimage", middleware.JwtAuth(), imagesApi.FreeImagesUploadView)

	router.GET("images", imagesApi.ImageListView)
	router.GET("image_names", imagesApi.ImageNameListView)

	router.PUT("images", middleware.JwtAuth(), imagesApi.ImageUpdateView)

	router.DELETE("images", middleware.JwtAuth(), imagesApi.ImageRemoveView)

}
