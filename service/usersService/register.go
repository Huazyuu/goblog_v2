package usersService

import (
	"backend/models/diverseType"
	"backend/repository/user_repo"
	"backend/utils"
	"errors"
)

func UserRegister(username, nickname, password string, role diverseType.Role, email string, ip string) error {
	user, err := user_repo.GetByUserName(username)
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
	user.Addr = utils.GetAddr(ip)

	err = user_repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
