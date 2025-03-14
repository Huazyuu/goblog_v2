package usersService

import (
	"backend/controller/req"
	"backend/global"

	"backend/repository/user_repo"
	"errors"
	"fmt"
)

func UserRemove(cr req.RemoveRequest) (string, error) {
	ulist, err := user_repo.GetUsersByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(ulist) == 0 {
		return "用户不存在", errors.New("用户不存在")
	}

	cnt, err := user_repo.DeleteUsers(ulist)
	if err != nil {
		global.Log.Error(err.Error())
		return "删除失败", err
	}
	return "删除成功 " + fmt.Sprintf("共删除 %d 个用户", cnt), err

}
