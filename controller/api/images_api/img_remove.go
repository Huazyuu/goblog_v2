package images_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"

	"backend/service/fileService"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := fileService.ImageRemoveById(cr)
	if err != nil {
		global.Log.Error(msg, err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
