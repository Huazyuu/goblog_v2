package advertsService

import (
	"backend/global"
	"backend/models/req"
	"backend/models/sqlmodels"
	"errors"
	"github.com/fatih/structs"
)

func AdvertUpdateService(id string, cr req.AdvertRequest) (string, error) {
	var ad sqlmodels.AdvertModel

	err := ad.GetByID(id)
	if err != nil {
		return "广告不存在", errors.New("广告不存在")
	}
	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	global.Log.Info(maps)
	err = ad.UpdateAdvert(maps)
	if err != nil {
		return "修改广告失败", errors.New("修改广告失败")
	}
	return "修改成功", nil
}
