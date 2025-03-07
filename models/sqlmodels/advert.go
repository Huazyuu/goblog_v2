package sqlmodels

import "backend/global"

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32;comment:广告标题" json:"title"` // 显示的标题
	Href   string `gorm:"comment:跳转链接" json:"href"`          // 跳转链接
	Images string `gorm:"comment:图片" json:"images"`          // 图片
	IsShow bool   `gorm:"comment:是否展示" json:"is_show"`       // 是否展示
}

func (*AdvertModel) TableName() string {
	return "advert"
}
func (a *AdvertModel) CreateAdvert() (err error) {
	return global.DB.Create(&a).Error
}

func (a *AdvertModel) GetByTitle(title string) error {
	return global.DB.Take(&a, "title = ?", title).Error
}
func (a *AdvertModel) GetByID(id any) error {
	return global.DB.Take(&a, "id = ?", id).Error
}

func (a *AdvertModel) UpdateAdvert(mapdata map[string]any) error {
	return global.DB.Model(&a).Updates(mapdata).Error
}
func (a *AdvertModel) GetAdvertsByIDList(idList []uint) (advertList []AdvertModel, err error) {
	err = global.DB.Find(&advertList, idList).Error
	return advertList, err
}
