package menu_api

import (
	"backend/global"
	"backend/models/req"
	"backend/models/res"
	"backend/service/menuService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr req.MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 获取菜单ID
	idStr := c.Param("id")
	menuID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res.FailWithMessage("无效的菜单ID", c)
		return
	}
	msg, err := menuService.UpdateMenu(uint(menuID), cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
