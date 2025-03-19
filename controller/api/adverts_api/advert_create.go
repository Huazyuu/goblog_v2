package adverts_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/advertsService"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertCreateView(c *gin.Context) {
	log := logStash.NewAction(c)
	title := "创建广告"
	var cr req.AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err)
		res.FailWithError(err, &cr, c)
		return
	}

	msg, err := advertsService.AdvertCreateService(cr)
	if err != nil {
		log.ErrItem(title, "创建广告错误", err)
		res.FailWithMessage(msg, c)
		return
	}

	log.InfoItem(title, "创建结果", map[string]interface{}{"title": cr.Title})

	res.OkWithMessage(msg, c)
}
