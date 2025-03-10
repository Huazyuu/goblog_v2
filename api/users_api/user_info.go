package users_api

import (
	"backend/middleware/jwt"
	"backend/models/res"
	"backend/repository/user_repo"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	user, err := user_repo.GetByID(claims.UserID)
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	res.OkWithData(user, c)
}
