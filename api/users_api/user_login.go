package users_api

import (
	"backend/models/res"
	"backend/service/usersService"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UsersApi) UserLoginView(c *gin.Context) {
	var cr LoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
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
