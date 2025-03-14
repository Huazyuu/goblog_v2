package advertsService

import (
	"backend/controller/req"
	"backend/global"

	"backend/repository/advert_repo"
	"errors"
	"fmt"
)

func AdvertsRemoveById(cr req.RemoveRequest) (string, error) {
	adlist, err := advert_repo.GetAdvertsByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(adlist) == 0 {
		return "广告不存在", errors.New("广告不存在")
	}
	cnt, err := advert_repo.DeleteAdverts(adlist)
	if err != nil {
		global.Log.Error(err.Error())
		return "删除失败", err
	}
	return "删除成功 " + fmt.Sprintf("共删除 %d 条广告", cnt), err
}
