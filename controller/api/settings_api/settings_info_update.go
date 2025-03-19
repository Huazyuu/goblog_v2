package settings_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"backend/service/settingsService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr req.SettingsUriRequest
	title := "更新某一项设置"
	log := logStash.NewAction(c)
	log.SetRequest(c)

	err := c.ShouldBindUri(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	resFailCode, err := settingsService.UpdateSettingsInfo(c, cr.Name)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithCode(resFailCode.(res.ErrorCode), c)
		global.Log.Error(err)
		return
	}
	err = settingsService.SetYaml()
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage("设置信息错误", c)
		global.Log.Error(err.Error())
		return
	}
	log.WarnItem(title, "更新成功", fmt.Sprintf("创建%s成功", cr.Name))
	res.OkWith(c)
}
