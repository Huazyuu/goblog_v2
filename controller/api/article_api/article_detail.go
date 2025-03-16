package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/repository/article_repo"
	"backend/service/articleService"
	"backend/service/redisService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr req.ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// keyword 和 title 一样,keyword能精准查找

	model, err := article_repo.GetArticleByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

func (ArticleApi) ArticleDetailByIDView(c *gin.Context) {
	var cr req.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	article, err := article_repo.GetArticleByID(cr.ID)
	redisService.NewArticleLook().Set(cr.ID)

	if err != nil {
		res.FailWithMessage("没有该文章", c)
		return
	}
	res.OkWithData(article, c)
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr req.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	resp, msg, err := articleService.ArticleDetailService(c, cr.ID)
	if err != nil {
		res.FailWithMessage(msg, c)
		global.Log.Error(err.Error())
		return
	}
	res.OkWithData(resp, c)
}
