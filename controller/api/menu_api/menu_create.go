package menu_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
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
	msg, err := menuService.MenuCreateService(cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
