package tag_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"time"
)

func GetByTitle(title string) (sqlmodels.TagModel, error) {
	var model sqlmodels.TagModel
	err := global.DB.Where("title = ?", title).Take(&model).Error
	return model, err
}
func IsExistTagByTitle(title string) (bool, error) {
	var model sqlmodels.TagModel
	err := global.DB.Where("title = ?", title).Take(&model).Error
	if err != nil {
		return false, err
	}
	return true, nil
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

func GetTagCreateTimes(tags []string) (map[string]string, error) {
	var tagModels []sqlmodels.TagModel
	if err := global.DB.
		Where("title IN ?", tags).
		Find(&tagModels).Error; err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, t := range tagModels {
		result[t.Title] = t.CreatedAt.Format(time.DateTime)
	}
	return result, nil
}
