package router

import (
	"backend/api"
)

func (router RouterGroup) UsersRouter() {
	usersApi := api.ApiGroupApp.UsersApi

	router.POST("users/register", usersApi.UserRegister)
	router.POST("users/login", usersApi.UserLoginView)

}
