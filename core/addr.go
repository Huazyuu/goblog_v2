package core

import (
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/sirupsen/logrus"
	"log"
)

func InitAddrDB() *geoip2.DBReader {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		log.Fatal("ip地址数据库加载失败", err)
	}
	logrus.Info("init addr db success")
	return db
}
