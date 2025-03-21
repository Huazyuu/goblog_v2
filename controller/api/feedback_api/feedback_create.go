package feedback_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/plugins/logStash"
	"github.com/gin-gonic/gin"
)

func (FeedbackApi) FeedBackCreateView(c *gin.Context) {
	title := "用户反馈"
	log := logStash.NewAction(c)
	var cr req.FeedBackCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "反馈错误", "绑定参数错误")
		res.FailWithError(err, &cr, c)
		return
	}
	var model sqlmodels.FeedbackModel
	err = global.DB.Take(&model, "email = ? and content = ?", cr.Email, cr.Content).Error
	if err == nil {
		res.FailWithMessage("存在相同留言", c)
		log.ErrItem(title, "反馈错误", "存在相同留言")
		return
	}
	global.DB.Create(&sqlmodels.FeedbackModel{
		Email:   cr.Email,
		Content: cr.Content,
	})
	res.OkWithMessage("反馈成功", c)
	log.InfoItem(title, "反馈成功", cr.Content)
}
