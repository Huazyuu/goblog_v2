package global

import (
	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.DBReader
)
