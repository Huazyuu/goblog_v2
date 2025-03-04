package core

import "gvb_server/global"

func InitCore() {
	global.Config = initConf()
	global.Log = initLogger()
	global.AddrDB = initAddrDB()
	global.DB = initGorm()
	global.Redis = initRedis()
	global.ESClient = initES()
	defer global.AddrDB.Close()
}
