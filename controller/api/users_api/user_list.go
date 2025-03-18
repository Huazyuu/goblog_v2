package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"

	"backend/service/usersService"
	"github.com/gin-gonic/gin"
)

func (UsersApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var page req.UserListRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, cnt, err := usersService.UsersList(claims, page)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("查询出错:"+err.Error(), c)
		return
	}
	res.OkWithList(list, cnt, c)
}
