package menu_banner_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"gorm.io/gorm"
)

func GetMenuBannersByIDList(menuIDs []uint) ([]sqlmodels.MenuBannerModel, error) {
	// Preload 是go struct名 不是sql表名
	var menuBanners []sqlmodels.MenuBannerModel
	err := global.DB.Preload("BannerModel").
		Where("menu_id in ?", menuIDs).
		Order("sort desc").
		Find(&menuBanners).Error
	return menuBanners, err
}
func GetMenuBannersByID(id uint) ([]sqlmodels.MenuBannerModel, error) {
	var menuBanners []sqlmodels.MenuBannerModel
	err := global.DB.Preload("BannerModel").
		Where("menu_id = ?", id).
		Order("sort desc").
		Find(&menuBanners).Error
	return menuBanners, err
}

func CreateMenuBanners(tx *gorm.DB, menuBanners []sqlmodels.MenuBannerModel) error {
	return tx.Create(&menuBanners).Error
}

func DeleteMenuBanners(tx *gorm.DB, menuID uint) error {
	return tx.Where("menu_id = ?", menuID).Delete(&sqlmodels.MenuBannerModel{}).Error
}

func ClearByMenuID(tx *gorm.DB, menuID uint) error {
	return tx.Where("menu_id = ?", menuID).
		Delete(&sqlmodels.MenuBannerModel{}).Error
}

func CreateBatch(tx *gorm.DB, menuBanners []sqlmodels.MenuBannerModel) error {
	return tx.Create(&menuBanners).Error
}
