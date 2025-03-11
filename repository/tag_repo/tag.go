package tag_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

func GetByTitle(title string) (sqlmodels.TagModel, error) {
	var model sqlmodels.TagModel
	err := global.DB.Where("title = ?", title).Take(&model).Error
	return model, err
}
func GetByID(id uint) (sqlmodels.TagModel, error) {
	var model sqlmodels.TagModel
	err := global.DB.Where(id).Take(&model).Error
	return model, err
}
func CreateTag(tag sqlmodels.TagModel) error {
	return global.DB.Create(&tag).Error
}
func UpdateTag(id uint, mapdata map[string]any) error {
	return global.DB.Model(&sqlmodels.TagModel{}).Where(id).Updates(mapdata).Error
}
func DeleteTags(tags []sqlmodels.TagModel) (count int64, err error) {
	res := global.DB.Delete(&tags)
	count = res.RowsAffected
	err = res.Error
	return count, err
}

func GetTagsByIDList(idList []uint) (tagLists []sqlmodels.TagModel, err error) {
	err = global.DB.Find(&tagLists, idList).Error
	return tagLists, err
}
