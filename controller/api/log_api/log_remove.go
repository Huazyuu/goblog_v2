package log_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/plugins/logStash"
	"backend/plugins/sync"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (LogApi) LogRemoveListView(c *gin.Context) {
	err := sync.SyncLogs()
	if err != nil {
		res.FailWithMessage("系统错误err:"+err.Error(), c)
		return
	}
	title := "日志记录"
	log := logStash.NewAction(c)
	var cr req.RemoveRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "删除记录", "删除失败")
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []sqlmodels.LogModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		log.ErrItem(title, "删除记录", "日志不存在")
		res.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&list)
	log.WarnItem(title, "删除记录", fmt.Sprintf("共删除 %d 个日志", count))
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)
}
