package utils

import (
	"backend/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

// IsPrivateIPAddr 是否内网地址
func IsPrivateIPAddr(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return true
	}
	// 192.168
	// 172.16 - 172.31
	// 10
	// 169.254
	return (ip4[0] == 192 && ip4[1] == 168) ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
		(ip4[0] == 10) ||
		(ip4[0] == 169 && ip4[1] == 254)
}
func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip)
	if IsPrivateIPAddr(parseIP) {
		return "内网地址"
	}
	record, err := global.AddrDB.City(net.ParseIP(ip))
	if err != nil {
		return "错误的地址"
	}
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}
	city := record.City.Names["zh-CN"]
	return fmt.Sprintf("%s-%s", province, city)
}
func GetAddrByGin(c *gin.Context) (ip, addr string) {
	ip = c.ClientIP()
	addr = GetAddr(ip)
	return ip, addr
}

// IsPublicIPAddr 是否公网地址
func IsPublicIPAddr(ip string) bool {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	ip4 := IP.To4()
	if ip4 == nil {
		return false
	}
	if !IP.IsPrivate() && !IP.IsLoopback() {
		return true
	}
	return false
}
