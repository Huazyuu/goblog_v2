package menu_repo

import (
	"backend/models/sqlmodels"
	"gorm.io/gorm"
)

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

func GetMenuByID(tx *gorm.DB, id uint) (sqlmodels.MenuModel, error) {
	var menu sqlmodels.MenuModel
	err := tx.First(&menu, id).Error
	return menu, err
}
