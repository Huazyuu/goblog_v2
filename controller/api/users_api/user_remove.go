package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/usersService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserRemoveView(c *gin.Context) {
	title := "删除用户"
	log := logStash.NewAction(c)

	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := usersService.UserRemove(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "删除成功", fmt.Sprintf("删除用户id:%T", cr.IDList))
	res.OkWithMessage(msg, c)
}
