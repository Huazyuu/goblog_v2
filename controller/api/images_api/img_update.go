package images_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/fileService"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr req.ImageNameUpdateRequest
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
