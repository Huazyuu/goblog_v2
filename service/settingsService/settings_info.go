package settingsService

import (
	"backend/config"
	"backend/controller/res"
	"backend/global"

	"errors"
	"github.com/gin-gonic/gin"
)

func GetSettingsInfo(name string) (any, bool) {
	switch name {
	case "email":
		return global.Config.Email, true
	case "qq":
		return global.Config.QQ, true
	case "qiniu":
		return global.Config.QiNiu, true
	case "jwt":
		return global.Config.Jwt, true
	case "chat_group":
		return global.Config.ChatGroup, true
	case "gaode":
		return global.Config.Gaode, true
	default:
		return "没有配置对应信息", false
	}
}
func UpdateSettingsInfo(c *gin.Context, name string) (any, error) {
	switch name {
	case "email":
		var info config.Email
		err := c.ShouldBindJSON(&info)
		if err != nil {
			return res.ArgumentError, err
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err := c.ShouldBindJSON(&info)
		if err != nil {
			return res.ArgumentError, err
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err := c.ShouldBindJSON(&info)
		if err != nil {
			return res.ArgumentError, err
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err := c.ShouldBindJSON(&info)
		if err != nil {
			return res.ArgumentError, err
		}
		global.Config.Jwt = info
	case "chat_group":
		var info config.ChatGroup
		err := c.ShouldBindJSON(&info)
		if err != nil {
			return res.ArgumentError, err
		}
		global.Config.ChatGroup = info
	case "gaode":
		var info config.Gaode
		err := c.ShouldBindJSON(&info)
		if err != nil {
			return res.ArgumentError, err
		}
		global.Config.Gaode = info
	default:
		return res.ArgumentError, errors.New("没有对应的配置文件")
	}
	return nil, nil
}
