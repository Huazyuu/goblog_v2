package main

import (
	"backend/core"
	"backend/flags"
	"backend/global"
	"backend/plugins/script/cron"
	"backend/plugins/sync"
	"backend/router"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化核心配置
	core.InitCore()
	// 定时任务
	cron.CronInit()
	// 解析命令行参数
	option := flags.Parse()
	if option.Run() {
		return
	}

	r := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Info("初始化路由成功")
	global.Log.Infof("服务运行在：http://%s/apiv1", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Log.Fatalf("服务器启动失败: %s", err.Error())
		}
	}()
	exitGracefully(srv)
}

func exitGracefully(srv *http.Server) {
	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Log.Info("接收到关闭信号，开始关闭流程...")

	// 创建同步操作的独立上下文（30秒超时）
	syncCtx, syncCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer syncCancel()

	// 执行数据同步
	global.Log.Info("开始同步文章数据到Elasticsearch...")
	if err := sync.SyncArticleData(syncCtx); err != nil {
		global.Log.Errorf("es数据同步失败: %v", err)
	} else {
		global.Log.Info("es数据同步完成")
	}
	global.Log.Info("开始同步文章数据到Mysql...")
	if err := sync.SyncCommentData(syncCtx); err != nil {
		global.Log.Errorf("mysql数据同步失败: %v", err)
	} else {
		global.Log.Info("mysql数据同步完成")
	}

	// 创建服务器关闭的上下文（15秒超时）
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	// 关闭HTTP服务器
	global.Log.Info("正在关闭HTTP服务器...")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		global.Log.Errorf("HTTP服务器关闭失败: %v", err)
	} else {
		global.Log.Info("HTTP服务器已正常关闭")
	}

	global.Log.Info("服务关闭流程完成")
}
