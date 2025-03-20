package log_api

import (
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/plugins/sync"
	"github.com/gin-gonic/gin"
)

func (LogApi) LogRefreshView(c *gin.Context) {
	// 手动保存redis数据到sql
	title := "日志记录"
	log := logStash.NewAction(c)

	err := sync.SyncLogs()
	if err != nil {
		log.ErrItem(title, "日志系统刷新错误", "同步数据失败")
		res.FailWithMessage("刷新出错", c)
		return
	}
	res.OkWithMessage("同步日志成功", c)
	log.WarnItem(title, "日志刷新成功", "同步数据成功")
}
