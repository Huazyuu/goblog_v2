package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/middleware/jwt"
	"backend/plugins/logStash"
	"fmt"

	"backend/service/usersService"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserBindEmailView(c *gin.Context) {
	title := "绑定邮箱"
	log := logStash.NewAction(c)
	log.SetRequest(c)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr req.BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}

	session := sessions.Default(c)
	msg, err := usersService.UserBindEmail(claims, session, cr.Email, cr.Password, cr.Code)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "绑定成功", fmt.Sprintf("用户%s 绑定 %s", claims.Username, cr.Email))
	res.OkWithMessage(msg, c)
}
