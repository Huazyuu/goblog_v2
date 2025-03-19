package tag_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/tagService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagCreateView(c *gin.Context) {
	var cr req.TagRequest
	title := "创建文章标签"
	log := logStash.NewAction(c)

	if err := c.ShouldBindJSON(&cr); err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := tagService.TagCreate(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.InfoItem(title, "创建成功", fmt.Sprintf("创建标签%s:", cr.Title))
	res.OkWithMessage(msg, c)
}
