package tagService

import (
	"backend/controller/req"
	"backend/repository/tag_repo"
	"github.com/fatih/structs"
)

func TagUpdate(tagID uint, cr req.TagRequest) (string, error) {
	_, err := tag_repo.GetByID(tagID)
	if err != nil {
		return "标签不存在", err
	}
	maps := structs.Map(&cr)
	err = tag_repo.UpdateTag(tagID, maps)
	if err != nil {
		return "更新标签失败", err
	}
	return "更新标签成功", nil
}
