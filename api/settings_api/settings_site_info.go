package settings_api

import (
	"backend/global"
	"backend/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsSiteInfoView 显示网站信息
func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
