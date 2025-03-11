package tagService

import (
	"backend/models/req"
	"backend/models/sqlmodels"
	"backend/repository/tag_repo"
	"fmt"
)

func TagCreate(cr req.TagRequest) (string, error) {
	_, err := tag_repo.GetByTitle(cr.Title)
	if err == nil {
		return "标签已存在", err
	}
	err = tag_repo.CreateTag(sqlmodels.TagModel{
		Title: cr.Title,
	})
	if err != nil {
		return "添加标签失败", err
	}
	return fmt.Sprintf("创建标签 %s 成功", cr.Title), nil
}
