package router

import (
	"backend/controller/api"
	"backend/middleware"
)

func (router RouterGroup) MenuRouter() {
	menu := api.ApiGroupApp.MenuApi
	router.POST("menus", middleware.JwtAdmin(), menu.MenuCreateView)

	router.GET("menus", menu.MenuListView)
	router.GET("menu_names", menu.MenuNameListView)
	router.GET("menus/:id", menu.MenuDetailByIDView)
	router.GET("menus/detail", menu.MenuDetailByPathView)

	router.PUT("menus/:id", middleware.JwtAdmin(), menu.MenuUpdateView)

	router.DELETE("menus", middleware.JwtAdmin(), menu.MenuRemoveView)

}
