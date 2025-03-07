package images_api

import (
	"backend/models/res"
	"backend/service/fileService"
	"github.com/gin-gonic/gin"
)

type ImageNameUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageNameUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := fileService.ImageUpdateNameByID(cr.ID, cr.Name)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
	return
}
