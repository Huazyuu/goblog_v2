package logStash

import (
	"backend/global"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"github.com/gin-gonic/gin"
)

// NewSuccessLogin 登录成功的日志
func NewSuccessLogin(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := splitToken(auth)
	jwyPayLoad := parseToken(token)
	saveLoginLog("用户登录", "登录成功", jwyPayLoad.UserID, jwyPayLoad.UserName, true, c)
}

// NewFailLogin 登录失败的日志
func NewFailLogin(title, content, userName string, c *gin.Context) {
	saveLoginLog(title, content, 0, userName, false, c)
}

func saveLoginLog(title string, content string, userID uint, userName string, status bool, c *gin.Context) {
	ip := c.ClientIP()
	addr := getAddr(ip)
	global.DB.Create(&sqlmodels.LogModel{
		IP:       ip,
		Addr:     addr,
		Title:    title,
		Content:  content,
		UserID:   userID,
		UserName: userName,
		Status:   status,
		Type:     diverseType.LoginType,
	})
}
