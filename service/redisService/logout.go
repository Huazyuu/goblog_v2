package redisService

import (
	"backend/global"
	"backend/utils"
	"strings"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	parts := strings.SplitN(token, " ", 2)
	tokenString := parts[1]
	err := global.Redis.Set(prefix+tokenString, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
