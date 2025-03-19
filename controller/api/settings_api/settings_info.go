package settings_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/settingsService"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr req.SettingsUriRequest

	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	info, ok := settingsService.GetSettingsInfo(cr.Name)
	if ok {
		res.OkWithData(info, c)
	} else {
		res.FailWithMessage(info.(string), c)
	}
}
