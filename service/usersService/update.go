package usersService

import (
	"backend/global"
	"backend/models/req"
	"backend/models/sqlmodels"
	"backend/utils"
	"errors"
	"github.com/fatih/structs"
	"strings"
)

func UserUpdateRole(cr req.UserRoleRequest) (string, error) {
	var u sqlmodels.UserModel
	if err := u.GetUserById(int(cr.UserID)); err != nil {
		return "用户不存在 id错误", err
	}
	err := u.UpdateUser(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	})
	if err != nil {
		global.Log.Error(err)
		return "修改权限失败", err
	}
	return "修改成功", nil
}
func UserUpdatePwd(id int, old, new string) (string, error) {
	var user sqlmodels.UserModel
	err := user.GetUserById(id)
	if err != nil {
		return "用户不存在 id错误", err
	}
	if !utils.CheckPwd(user.Password, old) {
		return "密码错误", errors.New("密码错误")
	}
	if err = user.UpdateUser(map[string]any{
		"password": utils.EncryptPwd(new),
	}); err != nil {
		return "修改密码失败", err
	}
	return "修改密码成功", nil
}

func UserUpdateInfo(id int, cr req.UserUpdateInfoRequest) (string, error) {
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}

	var user sqlmodels.UserModel
	if err := user.GetUserById(id); err != nil {
		return "用户不存在", err
	}

	if err := user.UpdateUser(newMaps); err != nil {
		return "修改失败", err
	}
	return "修改成功", nil
}
