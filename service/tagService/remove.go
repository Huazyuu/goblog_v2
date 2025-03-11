package tagService

import (
	"backend/global"
	"backend/models/req"
	"backend/repository/tag_repo"
	"errors"
	"fmt"
)

func TagRemove(cr req.RemoveRequest) (string, error) {
	tlist, err := tag_repo.GetTagsByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(tlist) == 0 {
		return "标签不存在", errors.New("标签不存在")
	}

	cnt, err := tag_repo.DeleteTags(tlist)
	if err != nil {
		global.Log.Error(err.Error())
		return "删除失败", err
	}
	return "删除成功 " + fmt.Sprintf("共删除 %d 个标签", cnt), err
}
