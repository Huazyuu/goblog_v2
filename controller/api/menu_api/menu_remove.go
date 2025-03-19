package menu_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/menuService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	title := "删除菜单"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := menuService.MenuRemove(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "删除成功", fmt.Sprintf("删除菜单%T", cr.IDList))
	res.OkWithMessage(msg, c)
}
