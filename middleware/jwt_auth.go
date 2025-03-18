package middleware

import (
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/diverseType"

	"backend/service/redisService"
	"github.com/gin-gonic/gin"
	"strings"
)

// JwtAuth 中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		tokenString := SplitToken(authHeader)
		if tokenString == "" {
			res.FailWithMessage("token格式错误", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			global.Log.Errorf("JWT 解析错误: %v", err)
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		isLoggedOut := redisService.CheckLogout(tokenString)
		if isLoggedOut {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		tokenString := SplitToken(authHeader)
		if tokenString == "" {
			res.FailWithMessage("token格式错误", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			global.Log.Errorf("JWT 解析错误: %v", err)
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		isLoggedOut := redisService.CheckLogout(tokenString)
		if isLoggedOut {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		if claims.Role != int(diverseType.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}

// JwtParamsAuth jwt在params附带 中间件
func JwtParamsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// global.Log.Info(c.Request.Header)
		authHeader := c.Query("Authorization")
		if authHeader == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		tokenString := SplitToken(authHeader)
		if tokenString == "" {
			res.FailWithMessage("token格式错误", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			global.Log.Errorf("JWT 解析错误: %v", err)
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		isLoggedOut := redisService.CheckLogout(tokenString)
		if isLoggedOut {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// SplitToken Authorization:bearer xxx.xxx.xxx
func SplitToken(token string) string {
	if token == "" {
		return ""
	}
	parts := strings.SplitN(token, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	tokenString := parts[1]
	return tokenString
}
