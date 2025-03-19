package adverts_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"

	"backend/service/advertsService"
	"github.com/gin-gonic/gin"
)

func (AdvertsApi) AdvertRemoveView(c *gin.Context) {
	log := logStash.NewAction(c)
	title := "删除广告"

	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		log.ErrItem(title, "参数绑定错误", err.Error())
		return
	}
	msg, err := advertsService.AdvertsRemoveById(cr)
	if err != nil {
		global.Log.Error(msg, err)
		res.FailWithMessage(msg, c)
		log.ErrItem(title, "删除广告错误", msg)
		return
	}
	log.WarnItem(title, "删除广告成功", msg)
	res.OkWithMessage(msg, c)
}
