package usersService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/repository/user_repo"
	"backend/utils"
	"errors"
	"github.com/fatih/structs"
	"strings"
)

func UserUpdateRole(cr req.UserRoleRequest) (string, error) {
	var u sqlmodels.UserModel
	u, err := user_repo.GetByID(cr.UserID)
	if err != nil {
		return "用户不存在 id错误", err
	}
	err = user_repo.UpdateUser(u.ID, map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	})
	if err != nil {
		global.Log.Error(err)
		return "修改权限失败", err
	}
	return "修改成功", nil
}
func UserUpdatePwd(id uint, old, new string) (string, error) {
	user, err := user_repo.GetByID(id)
	if err != nil {
		return "用户不存在 id错误", err
	}
	if !utils.CheckPwd(user.Password, old) {
		return "密码错误", errors.New("密码错误")
	}
	if err = user_repo.UpdateUser(user.ID, map[string]any{
		"password": utils.EncryptPwd(new),
	}); err != nil {
		return "修改密码失败", err
	}
	return "修改密码成功", nil
}

func UserUpdateInfo(id uint, cr req.UserUpdateInfoRequest) (string, error) {
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}

	var user sqlmodels.UserModel
	user, err := user_repo.GetByID(id)
	if err != nil {
		return "用户不存在", err
	}

	if err := user_repo.UpdateUser(user.ID, newMaps); err != nil {
		return "修改失败", err
	}
	return "修改成功", nil
}
