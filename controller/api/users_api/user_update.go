package users_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/plugins/logStash"
	"fmt"

	"backend/service/usersService"
	"github.com/gin-gonic/gin"
)

// UserUpdateRoleView 管理员修改用户权限 昵称
func (UsersApi) UserUpdateRoleView(c *gin.Context) {
	title := "修改权限"
	log := logStash.NewAction(c)

	var cr req.UserRoleRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	info, err := usersService.UserUpdateRole(cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		global.Log.Error(err)
		res.FailWithMessage(info, c)
		return
	}
	res.FailWithMessage(info, c)
	log.InfoItem(title, "修改成功", fmt.Sprintf("修改用户id:%s权限为:%s", cr.NickName, cr.Role.String()))
}

// UserUpdatePasswordView 修改密码
func (UsersApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr req.UserUpdatePwdRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := usersService.UserUpdatePwd(claims.UserID, cr.OldPwd, cr.Pwd)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.FailWithMessage(msg, c)
}

// UserUpdateInfoView 修改用户nickname sign link avatar
func (UsersApi) UserUpdateInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr req.UserUpdateInfoRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	msg, err := usersService.UserUpdateInfo(claims.UserID, cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.FailWithMessage(msg, c)

}
