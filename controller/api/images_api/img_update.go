package images_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/fileService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr req.ImageNameUpdateRequest
	title := "修改图片名称"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := fileService.ImageUpdateNameByID(cr.ID, cr.Name)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "更新图片成功", fmt.Sprintf("更新图片%d名称为%s", cr.ID, cr.Name))
	res.OkWithMessage(msg, c)
	return
}
