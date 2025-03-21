package utils

import "strings"

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
