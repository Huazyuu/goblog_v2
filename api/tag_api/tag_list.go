package tag_api

import (
	"backend/models/common"
	"backend/models/res"
	"backend/service/tagService"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr common.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, cnt, err := tagService.TagList(cr)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithList(list, cnt, c)
}

func (TagApi) TagNameListView(c *gin.Context) {
	// todo tag list_name api 需要查询es 找出对应tag
}
