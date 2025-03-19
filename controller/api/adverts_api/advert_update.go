package adverts_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"backend/service/advertsService"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertUpdateView(c *gin.Context) {
	log := logStash.NewAction(c)
	title := "更新广告"

	var cr req.AdvertRequest
	id := c.Param("id")
	err := c.ShouldBind(&cr)
	if err != nil {
		log.ErrItem(title, "参数错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := advertsService.AdvertUpdateService(id, cr)
	if err != nil {
		global.Log.Error(err)
		log.SetItemErr("更新错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "更新广告成功", msg)
	res.OkWithMessage(msg, c)
}
