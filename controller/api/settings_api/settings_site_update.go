package settings_api

import (
	"backend/config"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"backend/service/settingsService"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var info config.SiteInfo
	title := "更新网站信息"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&info)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = info
	err = settingsService.SetYaml()
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage("设置信息错误", c)
		global.Log.Error(err.Error())
		return
	}
	log.WarnItem(title, "更新成功", "更新网站信息成功")
	res.OkWithMessage("网站信息更新成功", c)
}
