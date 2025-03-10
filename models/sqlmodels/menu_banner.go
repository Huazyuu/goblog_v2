package sqlmodels

import (
	"backend/global"
	"gorm.io/gorm"
)

// MenuBannerModel 自定义菜单和背景图的连接表
type MenuBannerModel struct {
	MenuID      uint        `gorm:"comment:菜单的id" json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `gorm:"comment:banner图的id" json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `gorm:"size:10;comment:序号" json:"sort"`
}

func (*MenuBannerModel) TableName() string {
	return "menu_banner"
}
func (m *MenuBannerModel) CreateMenuBanner(list []MenuBannerModel) error {
	return global.DB.Create(&list).Error
}

// CreateMenuBannerTX 批量创建关联关系（事务版）
func (m *MenuBannerModel) CreateMenuBannerTX(tx *gorm.DB, list []MenuBannerModel) error {
	return tx.Create(&list).Error
}
