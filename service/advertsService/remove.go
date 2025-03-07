package advertsService

import (
	"backend/global"
	"backend/models/req"
	"backend/models/sqlmodels"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AdvertsRemoveById(cr req.RemoveRequest) (string, error) {

	blist, err := (&sqlmodels.AdvertModel{}).GetAdvertsByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(blist) == 0 {
		return "广告不存在", errors.New("blist")
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&blist).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err.Error())
		return "删除失败", err
	}
	return "删除成功 " + fmt.Sprintf("共删除 %d 条广告", len(blist)), err
}
