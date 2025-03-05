package settings_api

import (
	"backend/models/res"
	"backend/service/settingsService"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	info, ok := settingsService.GetSettingsInfo(cr.Name)
	if ok {
		res.OkWithData(info, c)
	} else {
		res.FailWithMessage("info", c)
	}
}
