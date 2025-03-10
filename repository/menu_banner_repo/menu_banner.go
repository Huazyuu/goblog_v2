package menu_banner_repo

import (
	"backend/models/sqlmodels"
	"gorm.io/gorm"
)

func CreateMenuBanners(tx *gorm.DB, menuBanners []sqlmodels.MenuBannerModel) error {
	return tx.Create(&menuBanners).Error
}

func DeleteMenuBanners(tx *gorm.DB, menuID uint) error {
	return tx.Where("menu_id = ?", menuID).Delete(&sqlmodels.MenuBannerModel{}).Error
}
