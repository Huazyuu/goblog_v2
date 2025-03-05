package router

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("7@dL#9fVx5*A2cS6m4%8yKp3!zB7Q0rT"))

func (router RouterGroup) UsersRouter() {
	router.Use(sessions.Sessions("sessionid", store))

	usersApi := api.ApiGroupApp.UsersApi
	router.POST("users/register", usersApi.UserRegisterView)
	router.POST("users/login", usersApi.UserLoginView)
	router.POST("users/logout", middleware.JwtAuth(), usersApi.UserLogoutView)
	router.POST("users/bind_email", middleware.JwtAuth(), usersApi.UserBindEmailView)

	router.GET("users", middleware.JwtAuth(), usersApi.UserListView)
}
