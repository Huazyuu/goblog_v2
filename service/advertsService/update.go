package advertsService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/repository/advert_repo"
	"errors"
	"github.com/fatih/structs"
	"strconv"
)

func AdvertUpdateService(id string, cr req.AdvertRequest) (string, error) {
	var ad sqlmodels.AdvertModel
	numUint64, _ := strconv.ParseUint(id, 10, 0)
	numUint := uint(numUint64)
	ad, err := advert_repo.GetByID(numUint)
	if err != nil {
		return "广告不存在", errors.New("广告不存在")
	}
	// 结构体转map的第三方包
	maps := structs.Map(&cr)
	global.Log.Info(maps)
	err = advert_repo.UpdateAdvert(ad.ID, maps)
	if err != nil {
		return "修改广告失败", errors.New("修改广告失败")
	}
	return "修改成功", nil
}
