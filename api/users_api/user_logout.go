package users_api

import (
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/res"
	"backend/service/usersService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	token := c.Request.Header.Get("Authorization")
	err := usersService.UserLogout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	global.Log.Info(fmt.Sprintf("用户 %s 注销登录", claims.Username))
	res.OkWithMessage("注销成功", c)
}
