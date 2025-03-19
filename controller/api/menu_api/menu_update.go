package menu_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"fmt"

	"backend/service/menuService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr req.MenuRequest
	title := "更新菜单"
	log := logStash.NewAction(c)

	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 获取菜单ID
	idStr := c.Param("id")
	menuID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.ErrItem(title, "无效的菜单ID", err.Error())
		res.FailWithMessage("无效的菜单ID", c)
		return
	}
	msg, err := menuService.UpdateMenu(uint(menuID), cr)
	if err != nil {
		global.Log.Error(err)
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "更新成功", fmt.Sprintf("创建菜单%s", cr.Title))
	res.OkWithMessage(msg, c)
}
