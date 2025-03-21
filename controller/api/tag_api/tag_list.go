package tag_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/tagService"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr req.PageInfo
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
	tagList, err := tagService.TagNameListService()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(tagList, c)
}
