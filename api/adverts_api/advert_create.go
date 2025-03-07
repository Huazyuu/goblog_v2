package adverts_api

import (
	"backend/models/req"
	"backend/models/res"
	"backend/service/advertsService"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertCreateView(c *gin.Context) {
	var cr req.AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := advertsService.AdvertCreateService(cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
