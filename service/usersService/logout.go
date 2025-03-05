package usersService

import (
	"backend/middleware/jwt"
	"backend/service/redisService"
	"time"
)

func UserLogout(claims *jwt.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redisService.Logout(token, diff)
}
