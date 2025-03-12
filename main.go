package main

import (
	"backend/core"
	"backend/flags"
	"backend/global"
	"backend/router"
)

func main() {
	core.InitCore()
	// todo flag
	// 命令行参数绑定
	option := flags.Parse()
	if option.Run() {
		return
	}

	// todo cron script
	r := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Info("init routers success")
	global.Log.Infof("server 运行在： http://%s/apiv1", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
