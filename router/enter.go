package router

import (
	"backend/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// todo log middleware
	// router.Use(middleware.LogMiddleWare())
	/*
		将 uploads 目录中的文件映射到一个特定的 URL 路径
		具体来说，它会将以 /uploads 开头的 HTTP 请求映射到本地文件系统中的 uploads 目录
		当客户端发送一个以 /uploads 开头的请求时，
		服务器会尝试从本地的 uploads 目录中查找对应的文件，并将其返回给客户端
	*/
	router.StaticFS("uploads", http.Dir("uploads"))
	// todo swagger
	// router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r := RouterGroup{router.Group("apiv1")}
	// routers
	r.SettingsRouter()
	r.ImagesRouter()
	r.UsersRouter()
	r.AdvertRouter()
	return router
}
