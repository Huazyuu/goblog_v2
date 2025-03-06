package usersService

import (
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/sqlmodels"
	"backend/plugins/email_plugin"
	"backend/utils"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"math/rand"
)

func UserBindEmail(claims *jwt.CustomClaims, session sessions.Session, email, password string, code *string) (string, error) {
	// 第一次请求
	var sendCode string
	if code == nil {
		sendCode = fmt.Sprintf("%4v", rand.Intn(10000))

		session.Set("valid_code", sendCode)
		session.Set("valid_email", email)
		err := session.Save()
		if err != nil {
			global.Log.Error(err)
			return "session错误", errors.New("session错误")
		}

		err = email_plugin.NewCode().Send(email, "你的验证码是 "+sendCode)
		if err != nil {
			global.Log.Error(err)
		}
		return "验证码已发送", nil
	}
	// 第二次请求 用户接收到code
	secondEmail := session.Get("valid_email")
	receiveCode := session.Get("valid_code")
	// 校验验证码
	if secondEmail != email {
		return "请输入第一次请求的邮箱", errors.New("请输入第一次请求的邮箱")
	}
	if receiveCode != *code {
		return "验证码错误", errors.New("验证码错误")
	}

	// 修改邮箱
	var userModel sqlmodels.UserModel
	err := userModel.GetUserById(int(claims.UserID))
	if err != nil {
		return "用户不存在", err
	}
	if len(password) < 6 {
		return "密码强度太低", errors.New("密码强度太低")
	}
	err = userModel.UpdateUser(map[string]any{
		"email":    email,
		"password": utils.EncryptPwd(password),
	})
	if err != nil {
		global.Log.Error(err)
		return "邮箱绑定失败", err
	}
	return "邮箱绑定成功", nil
}
