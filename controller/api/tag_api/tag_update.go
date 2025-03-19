package tag_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/tagService"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	title := "更新标签"
	log := logStash.NewAction(c)

	id := c.Param("id")
	var cr req.TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	uint64id, _ := strconv.ParseUint(id, 10, 64)
	uintid := uint(uint64id)
	msg, err := tagService.TagUpdate(uintid, cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "更新成功", fmt.Sprintf("创建标签%s:", cr.Title))
	res.OkWithMessage(msg, c)
}
