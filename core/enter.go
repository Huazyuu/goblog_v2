package core

import "backend/global"

func InitCore() {
	global.Config = InitConf("")
	global.Log = InitLogger()
	global.AddrDB = initAddrDB()
	global.DB = initGorm()
	global.Redis = initRedis()
	// todo es
	// global.ESClient = initES()
	defer global.AddrDB.Close()
}
