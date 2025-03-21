package feedback_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/plugins/logStash"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (FeedbackApi) FeedBackRemoveView(c *gin.Context) {
	title := "用户反馈"
	log := logStash.NewAction(c)
	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "删除错误", "绑定参数错误")
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []sqlmodels.FeedbackModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		log.ErrItem(title, "删除错误", "内容不存在")
		res.FailWithMessage("内容不存在", c)
		return
	}
	err = global.DB.Delete(&list).Error
	if err != nil {
		log.ErrItem(title, "删除错误", "系统错误")
		res.FailWithMessage("删除反馈内容失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除 %d 条反馈内容", count), c)
	log.WarnItem(title, "删除成功", fmt.Sprintf("共删除 %d 条反馈内容", count))
}
