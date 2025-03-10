package menu_api

import (
	"backend/global"
	"backend/models/req"
	"backend/models/res"
	"backend/service/menuService"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr req.MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := menuService.MenuCreate(cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
