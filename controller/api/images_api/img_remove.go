package images_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"fmt"

	"backend/service/fileService"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	title := "删除图片"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := fileService.ImageRemoveById(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		global.Log.Error(msg, err)
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "删除成功", fmt.Sprintf("删除图片%T", cr.IDList))
	res.OkWithMessage(msg, c)
}
