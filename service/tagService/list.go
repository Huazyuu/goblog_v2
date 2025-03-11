package tagService

import (
	"backend/models/common"
	"backend/models/sqlmodels"
)

func TagList(cr common.PageInfo) ([]sqlmodels.TagModel, int64, error) {
	list, cnt, err := common.ComList(sqlmodels.TagModel{}, common.Option{
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
