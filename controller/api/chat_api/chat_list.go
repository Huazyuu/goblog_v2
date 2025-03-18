package chat_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/models/sqlmodels"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ChatApi) ChatListView(c *gin.Context) {
	var cr req.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	cr.Sort = "created_at desc"
	list, count, _ := req.ComList(sqlmodels.ChatModel{IsGroup: true}, req.Option{
		PageInfo: cr,
		Likes:    []string{"content"},
	})
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]sqlmodels.ChatModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	res.OkWithList(data, count, c)
}
