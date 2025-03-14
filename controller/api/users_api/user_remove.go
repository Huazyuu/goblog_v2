package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/usersService"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserRemoveView(c *gin.Context) {
	var cr req.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := usersService.UserRemove(cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
