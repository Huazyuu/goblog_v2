package sqlmodels

import (
	"backend/global"
	"backend/models/diverseType"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	MODEL
	Path        string                `gorm:"comment:图片路径" json:"path"`                                      // 图片路径
	Hash        string                `gorm:"comment:图片的hash值" json:"hash"`                                  // 图片的hash值，用于判断重复图片
	Name        string                `gorm:"size:38;comment:图片名称" json:"name"`                              // 图片名称
	ImageType   diverseType.ImageType `gorm:"default:1;comment:图片的类型，本地还是七牛,1本地，2七牛，默认是1" json:"image_type"` // 图片的类型， 本地还是七牛
	MenusBanner []MenuBannerModel     `gorm:"foreignKey:BannerID" json:"-"`
}

func (b *BannerModel) TableName() string {
	return "banner"
}
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == diverseType.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path[1:])
		if err != nil {
			global.Log.Error(err)
			return nil
		}
	}
	return nil
}

func (b *BannerModel) CreateBanner() (err error) {
	return global.DB.Create(&b).Error
}

func (b *BannerModel) GetByPath(old string) error {
	return global.DB.Where("path = ?", old).Find(&b).Error
}
func (b *BannerModel) GetByID(id uint) error {
	return global.DB.Take(&b, "id = ?", id).Error
}

func (b *BannerModel) GetByHash(hash string) error {
	return global.DB.Take(&b, "hash = ?", hash).Error
}
func (b *BannerModel) UpdateBanner(mapdata map[string]any) error {
	return global.DB.Model(&b).Updates(mapdata).Error
}
func (b *BannerModel) GetBannersByIDList(idList []uint) (bannerList []BannerModel, err error) {
	err = global.DB.Find(&bannerList, idList).Error
	return bannerList, err
}
