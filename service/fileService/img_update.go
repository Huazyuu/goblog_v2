package fileService

import (
	"backend/repository/img_repo"
	"fmt"
)

func ImageUpdateNameByID(id uint, name string) (string, error) {
	_, err := img_repo.GetByID(id)
	if err != nil {
		return "文件不存在", err
	}
	err = img_repo.UpdateBanner(id, map[string]any{
		"name": name,
	})
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("修改图片(id:%d)%s 成功", id, name), nil
}
