package adverts_api

import (
	"backend/global"
	"backend/models/req"
	"backend/models/res"
	"backend/service/advertsService"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertUpdateView(c *gin.Context) {
	var cr req.AdvertRequest
	id := c.Param("id")
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := advertsService.AdvertUpdateService(id, cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
