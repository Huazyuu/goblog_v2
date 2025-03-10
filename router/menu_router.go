package router

import (
	"backend/api"
	"backend/middleware"
)

func (router RouterGroup) MenuRouter() {
	menu := api.ApiGroupApp.MenuApi
	router.POST("menus", middleware.JwtAdmin(), menu.MenuCreateView)

}
