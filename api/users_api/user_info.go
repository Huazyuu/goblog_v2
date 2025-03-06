package users_api

import (
	"backend/middleware/jwt"
	"backend/models/res"
	"backend/models/sqlmodels"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var u sqlmodels.UserModel
	err := u.GetUserById(int(claims.UserID))
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	res.OkWithData(u, c)
}
