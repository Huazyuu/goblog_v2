package logStash

import (
	"backend/global"
	"backend/utils"
	"fmt"
	"math"
	"net"
	"strings"
)

// formatBytes 格式化输出字节单位
func formatBytes(size int64) string {
	const unit = 1024
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	convertedSize := float64(size)
	unitIndex := 0
	for convertedSize >= unit && unitIndex < len(units)-1 {
		// 大于一个单位 并且单位列表有更大的单位
		convertedSize /= unit
		unitIndex++
	}

	// 四舍五入保留两位小数
	roundedSize := math.Round(convertedSize*100) / 100

	return fmt.Sprintf("%.2f %s", roundedSize, units[unitIndex])
}

func getAddr(ip string) (addr string) {
	if !utils.IsPublicIPAddr(ip) {
		return "内网地址"
	}
	cities, err := global.AddrDB.City(net.ParseIP(ip))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 国家
	country := cities.Country.Names["zh-CN"]
	// 城市
	city := cities.City.Names["zh-CN"]
	// 省份
	var subdivisions string
	if len(cities.Subdivisions) > 0 {
		subdivisions = cities.Subdivisions[0].Names["zh-CN"]
		return fmt.Sprintf("%s-%s", subdivisions, city)
	}
	if city != "" {
		return fmt.Sprintf("%s-%s", country, city)
	}
	return "未知地址"
}

// splitToken Authorization:bearer xxx.xxx.xxx
func splitToken(token string) string {
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
