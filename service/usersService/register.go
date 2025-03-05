package usersService

import (
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"backend/utils"
	"errors"
)

func UserRegister(username, nickname, password string, role diverseType.Role, email string, ip string) error {
	var user sqlmodels.UserModel
	err := user.ISUserExist(username)
	if err == nil {
		return errors.New("用户名已存在")
	}
	user.UserName = username
	user.NickName = nickname
	user.Password = utils.EncryptPwd(password)
	user.Role = role
	user.Avatar = "/uploads/avatar/default.png"
	user.Email = email
	user.IP = ip
	user.Addr = "内网地址"

	err = user.CreateUser(&user)
	if err != nil {
		return err
	}

	return nil
}
