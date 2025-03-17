package comment_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/service/commentService"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentListView(c *gin.Context) {
	var cr req.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, err := commentService.CommentListService(cr)
	if err != nil {
		res.OkWithMessage("查询失败 log: "+err.Error(), c)
		global.Log.Error(err)
		return
	}
	res.OkWithList(list, int64(len(list)), c)
}

func (CommentApi) CommentByArticleListView(c *gin.Context) {
	var cr req.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	resp, err := commentService.CommentByArticleListService(cr)
	if err != nil {
		res.OkWithMessage("查询失败 log: "+err.Error(), c)
		global.Log.Error(err)
		return
	}
	res.OkWithList(resp, int64(len(resp)), c)
}

func (CommentApi) CommentListInArticleView(c *gin.Context) {
	var cr req.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, cnt := commentService.CommentListInArticle(cr.ID)
	res.OkWithList(list, cnt, c)
}
