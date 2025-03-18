package logStash

import "github.com/dgrijalva/jwt-go/v4"

type JwtPayLoad struct {
	NickName string `json:"nick_name"`
	RoleID   uint   `json:"role"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"username"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func parseToken(token string) (jwtPayload *JwtPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if Token == nil || Token.Claims == nil {
		return nil
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	return &claims.JwtPayLoad
}
