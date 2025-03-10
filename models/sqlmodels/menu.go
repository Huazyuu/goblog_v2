package sqlmodels

import (
	"backend/global"
	"backend/models/diverseType"
	"gorm.io/gorm"
)

// MenuModel 菜单表  菜单的路径可以是 /path 也可以是路由别名
type MenuModel struct {
	MODEL
	Title        string            `gorm:"size:32;comment:菜单标题" json:"title"`                                                         // 标题
	Path         string            `gorm:"size:32;comment:菜单路径" json:"path"`                                                          // 路径
	Slogan       string            `gorm:"size:64;comment:slogan" json:"slogan"`                                                      // slogan
	Abstract     diverseType.Array `gorm:"comment:简介，按照换行去切割为数组" gorm:"type:string" json:"abstract"`                                  // 简介
	AbstractTime int               `gorm:"comment:简介的切换时间" json:"abstract_time"`                                                      // 简介的切换时间
	Banners      []BannerModel     `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` // 菜单的图片列表
	BannerTime   int               `gorm:"comment:banner图的切换时间" json:"banner_time"`                                                   // 菜单图片的切换时间 为 0 表示不切换
	Sort         int               `gorm:"size:10;comment:顺序" json:"sort"`                                                            // 菜单的顺序
}

func (*MenuModel) TableName() string {
	return "menu"
}
func (m *MenuModel) CreateMenuByParam(menu *MenuModel) error {
	return global.DB.Create(&MenuModel{
		Title:        menu.Title,
		Path:         menu.Path,
		Slogan:       menu.Slogan,
		Abstract:     menu.Abstract,
		AbstractTime: menu.AbstractTime,
		BannerTime:   menu.BannerTime,
		Sort:         menu.Sort,
	}).Error
}
func (m *MenuModel) CreateMenu() error {
	return global.DB.Create(m).Error
}

// IsDuplicate 使用事务进行重复检查
func (m MenuModel) IsDuplicate(tx *gorm.DB, title, path string) (int64, error) {
	var count int64
	err := tx.Model(&MenuModel{}).
		Where("title = ? OR path = ?", title, path).
		Count(&count).Error
	return count, err
}
