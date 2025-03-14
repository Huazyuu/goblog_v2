package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/models/esmodels"

	"backend/service/articleService"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr req.ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, msg, err := articleService.ArticleListService(c, cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithList(list, int64(len(list)), c)
}

type ArticleIDTitleListResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (ArticleApi) ArticleIDTitleListView(c *gin.Context) {
	var article esmodels.ArticleModel
	result, err := global.ESClient.Search(article.Index()).
		Query(elastic.NewMatchAllQuery()).Source(`{"_source": ["title"]}`). // 只返回文档中的 title 字段，而忽略其他字段
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}
	var respList []ArticleIDTitleListResponse
	for _, hit := range result.Hits.Hits {
		json.Unmarshal(hit.Source, &article)
		respList = append(respList, ArticleIDTitleListResponse{
			ID:    hit.Id,
			Title: article.Title,
		})
	}
	res.OkWithData(respList, c)
}

func (ArticleApi) ArticleCategoryListView(c *gin.Context) {
	list, msg, err := articleService.ArticleCategoryListService()
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithData(list, c)
}
