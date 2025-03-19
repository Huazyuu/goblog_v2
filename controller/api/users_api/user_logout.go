package users_api

import (
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/plugins/logStash"

	"backend/service/usersService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserLogoutView(c *gin.Context) {
	title := "用户登出"
	log := logStash.NewAction(c)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	token := c.Request.Header.Get("Authorization")
	err := usersService.UserLogout(claims, token)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	global.Log.Info(fmt.Sprintf("用户 %s 注销登录", claims.Username))
	log.WarnItem(title, "注销登录成功", fmt.Sprintf("用户 %s 注销登录", claims.Username))
	res.OkWithMessage("注销成功", c)
}
