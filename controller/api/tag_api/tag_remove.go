package tag_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/tagService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	title := "删除标签"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := tagService.TagRemove(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "删除成功", fmt.Sprintf("创建标签id:%T", cr.IDList))
	res.OkWithMessage(msg, c)
}
