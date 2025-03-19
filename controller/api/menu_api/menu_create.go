package menu_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"backend/service/menuService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr req.MenuRequest
	err := c.ShouldBindJSON(&cr)
	title := "创建菜单"
	log := logStash.NewAction(c)

	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := menuService.MenuCreateService(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	log.InfoItem(title, "创建成功", fmt.Sprintf("创建菜单:%s", cr.Title))
	res.OkWithMessage(msg, c)
}
