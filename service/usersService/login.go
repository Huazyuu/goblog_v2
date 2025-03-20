package usersService

import (
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/sqlmodels"
	"backend/plugins/logStash"
	"backend/repository/loginData_repo"
	"backend/repository/user_repo"
	"backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context, username, password string) (string, error) {
	userModel, err := user_repo.GetByUserName(username)
	title := "用户登录"
	if err != nil {
		logStash.NewFailLogin(title, "用户名不存在", username, c)
		global.Log.Warn("用户名不存在")
		return "", errors.New("用户不存在")
	}
	isCheck := utils.CheckPwd(userModel.Password, password)
	if !isCheck {
		logStash.NewFailLogin(title, "密码错误", username, c)
		global.Log.Warn("用户名密码错误")
		return "", errors.New("用户名密码错误")
	}
	// 登录成功
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		Username: userModel.UserName,
	})

	c.Request.Header.Set("Authorization", "bearer "+token)

	global.Log.Info(userModel)
	ip, addr := utils.GetAddrByGin(c)
	err = loginData_repo.CreateLoginData(sqlmodels.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.UserName,
		Token:     token,
		Addr:      addr,
		LoginType: 0,
		Device:    "",
	})
	if err != nil {
		logStash.NewFailLogin(title, "系统错误", username, c)
		return "", err
	}
	logStash.NewSuccessLogin(c)
	return token, nil
}
