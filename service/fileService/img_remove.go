package fileService

import (
	"backend/global"
	"backend/models/req"
	"backend/repository/img_repo"
	"errors"
	"fmt"
)

func ImageRemoveById(cr req.RemoveRequest) (string, error) {
	blist, err := img_repo.GetBannersByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(blist) == 0 {
		return "图片不存在", errors.New("blist")
	}
	count, err := img_repo.DeleteImgs(blist)
	if err != nil {
		global.Log.Error(err.Error())
		return "删除失败", err
	}
	return "删除成功 " + fmt.Sprintf("共删除 %d 张图片", count), err
}
