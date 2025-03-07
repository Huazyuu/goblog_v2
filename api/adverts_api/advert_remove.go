package adverts_api

import (
	"backend/global"
	"backend/models/req"
	"backend/models/res"
	"backend/service/advertsService"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := advertsService.AdvertsRemoveById(cr)
	if err != nil {
		global.Log.Error(msg, err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
