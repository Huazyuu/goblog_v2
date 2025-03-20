package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/service/usersService"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserLoginView(c *gin.Context) {
	var cr req.LoginRequest
	title := "用户登录"

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		logStash.NewFailLogin(title, "参数错误", "", c)
		res.FailWithError(err, &cr, c)
		return
	}

	token, err := usersService.UserLogin(c, cr.UserName, cr.Password)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(token, c)
}
