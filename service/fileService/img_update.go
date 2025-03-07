package fileService

import (
	"backend/models/sqlmodels"
	"fmt"
)

func ImageUpdateNameByID(id uint, name string) (string, error) {
	var imageModel sqlmodels.BannerModel
	err := imageModel.GetByID(id)
	if err != nil {
		return "文件不存在", err
	}
	err = imageModel.UpdateBanner(map[string]any{
		"name": name,
	})
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("修改图片(id:%d)%s 成功", id, name), nil
}
