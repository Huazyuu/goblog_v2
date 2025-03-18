package flags

import "flag"

type Option struct {
	DB   bool   // 创建数据库
	Load string // 导入数据库文件
	Dump bool   // 导出数据库

	EsCreate bool   // 创建索引
	ESDump   bool   // 导出es索引
	ESLoad   string // 导入es索引

	AvatarCreate bool
}

// Parse 解析命令行参数
func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "dbcreate", false, "初始化数据库")
	flag.StringVar(&option.Load, "dbload", "", "导入sql数据库")
	flag.BoolVar(&option.Dump, "dbdump", false, "导出sql数据库")

	flag.BoolVar(&option.EsCreate, "escreate", false, "创建索引")
	flag.StringVar(&option.ESLoad, "esload", "", "导入es索引")
	flag.BoolVar(&option.ESDump, "esdump", false, "导出es索引")

	flag.BoolVar(&option.AvatarCreate, "avatar", false, "生成头像")

	flag.Parse()
	return option
}

// Run 根据命令执行不同的函数
func (option Option) Run() bool {
	switch {
	case option.DB:
		dbCreate()
		return true
	case option.Dump:
		dbDump()
		return true
	case option.EsCreate:
		esCreate()
		return true
	case option.ESDump:
		esDump()
		return true
	case option.AvatarCreate:
		avatarCreate()
		return true
	}

	return false
}
