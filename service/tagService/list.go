package tagService

import (
	"backend/controller/req"
	"backend/models/sqlmodels"
)

func TagList(cr req.PageInfo) ([]sqlmodels.TagModel, int64, error) {
	list, cnt, err := req.ComList(sqlmodels.TagModel{}, req.Option{
		PageInfo: cr,
	})
	if err != nil {
		return nil, 0, err
	}
	return list, cnt, nil
}
func TagNameList() {
	// todo tag list_name service 需要查询es 找出对应tag
}
