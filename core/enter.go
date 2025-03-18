package core

import (
	"backend/global"
)

func InitCore() {
	global.Config = InitConf()
	global.Log = InitLogger()
	global.AddrDB = InitAddrDB()
	defer global.AddrDB.Close()
	global.DB = InitGorm()
	global.Redis = InitRedis()
	global.ESClient = InitES()
}
