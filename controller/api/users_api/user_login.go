package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/usersService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserLoginView(c *gin.Context) {
	var cr req.LoginRequest
	title := "用户登录"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}

	token, err := usersService.UserLogin(c, cr.UserName, cr.Password)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}
	log.InfoItem(title, "登陆成功", fmt.Sprintf("用户 %s 登录", cr.UserName))
	res.OkWithData(token, c)
}
