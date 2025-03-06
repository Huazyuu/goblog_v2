package users_api

import (
	"backend/middleware/jwt"
	"backend/models/req"
	"backend/models/res"
	"backend/service/usersService"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr req.BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	session := sessions.Default(c)
	msg, err := usersService.UserBindEmail(claims, session, cr.Email, cr.Password, cr.Code)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
