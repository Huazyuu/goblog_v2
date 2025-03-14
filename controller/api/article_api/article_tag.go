package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/articleService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleTagListView(c *gin.Context) {
	var cr req.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	response, total, err := articleService.GetArticleTags(cr.Page, cr.Limit)
	if err != nil {
		res.FailWithMessage("获取标签列表失败", c)
		return
	}
	res.OkWithList(response, total, c)
}
