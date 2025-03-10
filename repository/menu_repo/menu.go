package menu_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"errors"
	"gorm.io/gorm"
)

func GetMenuByID(id uint) (sqlmodels.MenuModel, error) {
	var menu sqlmodels.MenuModel
	err := global.DB.Take(&menu, id).Error
	return menu, err
}
func GetMenuByPath(path string) (sqlmodels.MenuModel, error) {
	var menu sqlmodels.MenuModel
	err := global.DB.Where("path = ?", path).Take(&menu).Error
	return menu, err
}

func GetMenuList() ([]sqlmodels.MenuModel, error) {
	var menus []sqlmodels.MenuModel
	err := global.DB.Order("sort desc").Find(&menus).Error
	return menus, err
}

func GetMenuIDs() ([]uint, error) {
	var ids []uint
	err := global.DB.Model(&sqlmodels.MenuModel{}).Order("sort desc").Pluck("id", &ids).Error
	return ids, err
}

func GetMenusByIDList(idList []uint) (menuList []sqlmodels.MenuModel, err error) {
	err = global.DB.Find(&menuList, idList).Error
	return menuList, err
}

func DeleteMenus(menuList []sqlmodels.MenuModel) (count int64, err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		for _, menu := range menuList {
			if err = tx.Where("menu_id = ?", menu.ID).Delete(&sqlmodels.MenuBannerModel{}).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					global.Log.Error(err)
					return err
				}
				// 如果是记录未找到的错误，忽略并继续
				err = nil
			}
		}
		// 删除主表记录
		res := tx.Delete(&menuList)
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

func UpdateMenu(tx *gorm.DB, id uint, updateMap map[string]interface{}) error {
	return tx.Model(&sqlmodels.MenuModel{}).
		Where("id = ?", id).
		Updates(updateMap).Error
}

func CheckMenuDuplicate(tx *gorm.DB, title, path string) (bool, error) {
	var count int64
	err := tx.Model(&sqlmodels.MenuModel{}).
		Where("title = ? OR path = ?", title, path).
		Count(&count).Error
	return count > 0, err
}

func CreateMenu(tx *gorm.DB, menu *sqlmodels.MenuModel) error {
	return tx.Create(menu).Error
}
