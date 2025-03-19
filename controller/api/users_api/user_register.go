package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"

	"backend/service/usersService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserRegisterView(c *gin.Context) {
	title := "用户注册"
	log := logStash.NewAction(c)

	var cr req.UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	err := usersService.UserRegister(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	log.InfoItem(title, "注册成功", fmt.Sprintf("用户%s注册,昵称:%s", cr.UserName, cr.NickName))
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
}
