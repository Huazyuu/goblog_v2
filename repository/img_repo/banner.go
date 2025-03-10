package img_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"gorm.io/gorm"
)

func CreateBanner(model sqlmodels.BannerModel) (err error) {
	return global.DB.Create(&model).Error
}

func GetByPath(old string) (sqlmodels.BannerModel, error) {
	var model sqlmodels.BannerModel
	err := global.DB.Where("path = ?", old).Take(&model).Error
	return model, err
}

func GetByID(id uint) (sqlmodels.BannerModel, error) {
	var model sqlmodels.BannerModel
	err := global.DB.Where("id = ?", id).Take(&model).Error
	return model, err
}

func GetByHash(hash string) (sqlmodels.BannerModel, error) {
	var model sqlmodels.BannerModel
	err := global.DB.Where("hash = ?", hash).Take(&model).Error
	return model, err
}

func UpdateBanner(id uint, mapdata map[string]any) error {
	return global.DB.Model(&sqlmodels.BannerModel{}).Where(id).Updates(mapdata).Error
}

func GetBannersByIDList(idList []uint) (bannerList []sqlmodels.BannerModel, err error) {
	err = global.DB.Find(&bannerList, idList).Error
	return bannerList, err
}

func DeleteImgs(imglist []sqlmodels.BannerModel) (count int64, err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 遍历要删除的 Banner 记录
		for _, banner := range imglist {
			// 删除外键关联的记录
			if err = tx.Where("banner_id = ?", banner.ID).Delete(&sqlmodels.MenuBannerModel{}).Error; err != nil {
				global.Log.Error(err)
				return err
			}
		}
		// 删除主表记录
		res := tx.Delete(&imglist)
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

/*================ 事务相关 ================*/

func BannerExists(tx *gorm.DB, id uint) (bool, error) {
	var count int64
	err := tx.Model(&sqlmodels.BannerModel{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func GetInvalidBannerIDs(tx *gorm.DB, ids []uint) ([]uint, error) {
	var validIDs []uint
	err := tx.Model(&sqlmodels.BannerModel{}).
		Where("id IN (?)", ids).
		Pluck("id", &validIDs).Error
	// Pluck 是一个用于从数据库中批量提取单个列值的方法
	if err != nil {
		return nil, err
	}

	invalidIDs := make([]uint, 0)
	for _, id := range ids {
		if !contains(validIDs, id) {
			invalidIDs = append(invalidIDs, id)
		}
	}
	return invalidIDs, nil
}

func contains(slice []uint, target uint) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
