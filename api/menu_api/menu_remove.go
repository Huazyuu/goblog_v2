package menu_api

import (
	"backend/models/req"
	"backend/models/res"
	"backend/service/menuService"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := menuService.MenuRemove(cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
