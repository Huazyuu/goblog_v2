package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// EncryptPwd hash密码
func EncryptPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// CheckPwd 验证密码  hash之后的密码  输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
