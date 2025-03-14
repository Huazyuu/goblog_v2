package settings_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/service/settingsService"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr req.SettingsUriRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	resFailCode, err := settingsService.UpdateSettingsInfo(c, cr.Name)
	if err != nil {
		res.FailWithCode(resFailCode.(res.ErrorCode), c)
		global.Log.Error(err)
		return
	}
	err = settingsService.SetYaml()
	if err != nil {
		res.FailWithMessage("设置信息错误", c)
		global.Log.Error(err.Error())
		return
	}
	res.OkWith(c)
}
