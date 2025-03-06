package usersService

import (
	"backend/global"
	"backend/models/req"
	"backend/models/sqlmodels"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func UserRemove(cr req.RemoveRequest) (string, error) {
	ulist, err := (&sqlmodels.UserModel{}).GetUsersByIDList(cr.IDList)
	if err != nil {
		global.Log.Error(err.Error())
		return "查找出错", err
	}
	if len(ulist) == 0 {
		return "用户不存在", errors.New("用户不存在")
	}

	// 事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&ulist).Error
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
	return "删除成功 " + fmt.Sprintf("共删除 %d 个用户", len(ulist)), err

}
