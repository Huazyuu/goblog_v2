package users_api

import (
	"backend/models/req"
	"backend/models/res"
	"backend/service/usersService"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserLoginView(c *gin.Context) {
	var cr req.LoginRequest
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
