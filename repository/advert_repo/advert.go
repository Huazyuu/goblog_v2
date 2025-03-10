package advert_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

func CreateAdvert(advert sqlmodels.AdvertModel) error {
	return global.DB.Create(&advert).Error
}

func GetByTitle(title string) (sqlmodels.AdvertModel, error) {
	var model sqlmodels.AdvertModel
	err := global.DB.Where("title = ?", title).Take(&model).Error
	return model, err
}

func GetByID(id uint) (sqlmodels.AdvertModel, error) {
	var model sqlmodels.AdvertModel
	err := global.DB.Where("id = ?", id).Take(&model).Error
	return model, err
}

func UpdateAdvert(id uint, mapdata map[string]any) error {
	return global.DB.Model(&sqlmodels.AdvertModel{}).Where(id).Updates(mapdata).Error
}

func GetAdvertsByIDList(idList []uint) (adList []sqlmodels.AdvertModel, err error) {
	err = global.DB.Find(&adList, idList).Error
	return adList, err
}

func DeleteAdverts(ads []sqlmodels.AdvertModel) (count int64, err error) {
	res := global.DB.Delete(&ads)
	count = res.RowsAffected
	err = res.Error
	return count, err
}
