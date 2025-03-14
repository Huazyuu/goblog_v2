package menuService

import (
	"backend/controller/req"
	"backend/global"
	"backend/repository/menu_repo"
	"errors"
	"fmt"
)

func MenuRemove(cr req.RemoveRequest) (string, error) {
	mlist, err := menu_repo.GetMenusByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(mlist) == 0 {
		return "菜单不存在", errors.New("菜单不存在")
	}

	cnt, err := menu_repo.DeleteMenus(mlist)
	if err != nil {
		global.Log.Error(err.Error())
		return "删除失败", err
	}
	return "删除成功 " + fmt.Sprintf("共删除 %d 条菜单", cnt), err
}
