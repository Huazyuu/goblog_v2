package settings_api

import (
	"backend/config"
	"backend/global"
	"backend/models/res"
	"backend/service/settingsService"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = info
	err = settingsService.SetYaml()
	if err != nil {
		res.FailWithError(err, "设置信息错误", c)
		global.Log.Error(err.Error())
		return
	}
	res.OkWithMessage("网站信息更新成功", c)
}
