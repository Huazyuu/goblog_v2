package core

import "backend/global"

func InitCore() {
	global.Config = initConf()
	global.Log = initLogger()
	global.AddrDB = initAddrDB()
	global.DB = initGorm()
	global.Redis = initRedis()
	// todo es
	// global.ESClient = initES()
	defer global.AddrDB.Close()
}
