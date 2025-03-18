package flags

import (
	"backend/global"
	"backend/models/sqlmodels"
	"backend/plugins/logStash"
)

func dbCreate() {
	var err error
	// MenuModel 和 Banners 之间存在多对多关系，并且使用 MenuBannerModel 作为连接表来记录这种关系
	global.DB.SetupJoinTable(&sqlmodels.MenuModel{}, "Banners", &sqlmodels.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&sqlmodels.BannerModel{},
			&sqlmodels.TagModel{},
			&sqlmodels.MessageModel{},
			&sqlmodels.AdvertModel{},
			&sqlmodels.UserModel{},
			&sqlmodels.CommentModel{},
			&sqlmodels.CollectModel{},
			&sqlmodels.MenuModel{},
			&sqlmodels.MenuBannerModel{},
			&sqlmodels.LoginDataModel{},
			&sqlmodels.ChatModel{},
			&sqlmodels.FeedbackModel{},
			&logStash.LogModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
