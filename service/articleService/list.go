package articleService

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware"
	"backend/middleware/jwt"
	"backend/models/esmodels"
	"backend/repository/article_repo"
	"backend/service/redisService"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
	"time"
)

func ArticleListService(c *gin.Context, cr req.ArticleSearchRequest) (list []esmodels.ArticleModel, msg string, err error) {
	// 是否显示我收藏的文章
	boolSearch := elastic.NewBoolQuery()
	if cr.IsUser {
		authHeader := c.GetHeader("Authorization")
		tokenString := middleware.SplitToken(authHeader)
		if tokenString == "" {
			res.FailWithMessage("token格式错误", c)
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err == nil && !redisService.CheckLogout(tokenString) {
			boolSearch.Must(elastic.NewTermsQuery("user_id", claims.UserID))
		}
	}

	if cr.Date != "" {
		date, err := time.Parse("2006-01-02", cr.Date)
		if err == nil {
			boolSearch.Must(elastic.NewRangeQuery("created_at").
				Gte(date.Format("2006-01-02") + " 00:00:00").
				Lte(date.Format("2006-01-02") + " 23:59:59"))
		}
	}

	list, _, err = article_repo.GetArticleList(article_repo.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
		Query:    boolSearch,
		Category: cr.Category,
	})
	if err != nil {
		global.Log.Error(err)
		return nil, "查询失败", err
	}
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		// 空对象返回数组
		list = make([]esmodels.ArticleModel, 0)
		return list, "查询成功", err
	}
	return list, "查询成功", err
}

type CategoryResponse struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

func ArticleCategoryListService() ([]CategoryResponse, string, error) {
	type termAggResult struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	agg := elastic.NewTermsAggregation().Field("category")
	res, err := global.ESClient.Search(esmodels.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).Aggregation("categories", agg).
		Size(0).Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return nil, "查询失败", err
	}

	byteData := res.Aggregations["categories"]
	var resp termAggResult
	_ = json.Unmarshal(byteData, &resp)
	var respList = make([]CategoryResponse, 0)
	for _, bucket := range resp.Buckets {
		respList = append(respList, CategoryResponse{
			Category: bucket.Key,
			Count:    bucket.DocCount,
		})
	}
	return respList, "查询成功", nil
}
