package user_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"errors"
	"gorm.io/gorm"
)

func CreateUser(user sqlmodels.UserModel) error {
	return global.DB.Create(&user).Error
}

func UpdateUser(id uint, mapdata map[string]any) error {
	return global.DB.Model(&sqlmodels.UserModel{}).Where(id).Updates(mapdata).Error
}

func GetByID(id uint) (sqlmodels.UserModel, error) {
	var model sqlmodels.UserModel
	err := global.DB.Where("id = ?", id).Take(&model).Error
	return model, err
}

func GetByUserName(username string) (sqlmodels.UserModel, error) {
	var model sqlmodels.UserModel
	err := global.DB.Where("user_name = ?", username).Take(&model).Error
	return model, err
}

func GetUsersByIDList(idList []uint) (usersList []sqlmodels.UserModel, err error) {
	err = global.DB.Find(&usersList, idList).Error
	return usersList, err
}

func GetAvatarByID(userID uint) (avatar string, err error) {
	err = global.DB.Model(sqlmodels.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	return avatar, err
}

func DeleteUsers(userList []sqlmodels.UserModel) (count int64, err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 遍历要删除的 User 记录
		for _, user := range userList {
			// 删除外键关联的记录 collect login message comment
			// 删除 CollectModel 记录
			if err = tx.Where("user_id = ?", user.ID).Delete(&sqlmodels.CollectModel{}).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					global.Log.Error(err)
					return err
				}
				// 如果是记录未找到的错误，忽略并继续
				err = nil
			}
			// 删除 LoginModel 记录
			if err = tx.Where("user_id = ?", user.ID).Delete(&sqlmodels.LoginDataModel{}).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					global.Log.Error(err)
					return err
				}
				// 如果是记录未找到的错误，忽略并继续
				err = nil
			}
			// 删除 MessageModel 记录
			if err = tx.Where("send_user_id = ? or rev_user_id = ?", user.ID, user.ID).Delete(&sqlmodels.MessageModel{}).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					global.Log.Error(err)
					return err
				}
				// 如果是记录未找到的错误，忽略并继续
				err = nil
			}
			// 删除 CommentModel 记录
			if err = tx.Where("user_id = ?", user.ID).Delete(&sqlmodels.CommentModel{}).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					global.Log.Error(err)
					return err
				}
				// 如果是记录未找到的错误，忽略并继续
				err = nil
			}
		}
		// 删除主表记录
		res := tx.Delete(&userList)
		err = res.Error
		count = res.RowsAffected
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
