package article_repo

import (
	"backend/global"
	"backend/models/esmodels"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

type TagAggregation struct {
	Tag           string   `json:"key"`
	DocCount      int64    `json:"doc_count"`
	ArticleIDList []string `json:"articles"`
}

type TagsAggregationResponse struct {
	Buckets []struct {
		Key      string `json:"key"`
		DocCount int64  `json:"doc_count"`
		Articles struct {
			Buckets []struct {
				Key string `json:"key"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

// GetTagAggregations 函数用于从 Elasticsearch 中获取文章标签的聚合信息，并支持分页处理
func GetTagAggregations(page, limit int) ([]TagAggregation, int64, error) {
	// 创建一个基数聚合对象，用于统计 tags 字段的唯一值数量
	countAgg := elastic.NewCardinalityAggregation().Field("tags")
	countResult, err := global.ESClient.
		Search(esmodels.ArticleModel{}.Index()).
		Aggregation("tags", countAgg).
		Size(0).
		Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	countAggResult, _ := countResult.Aggregations.Cardinality("tags")
	total := int64(*countAggResult.Value)

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	// 聚合查询
	// 创建一个词条聚合对象，按 tags 字段对文档进行分组
	agg := elastic.NewTermsAggregation().Field("tags")
	// 创建一个子聚合对象，按 keyword 字段对每个标签分组内的文档进行进一步分组
	subAgg := elastic.NewTermsAggregation().Field("keyword")
	// 创建一个桶排序聚合对象，用于实现分页功能，从 offset 位置开始，取 limit 条记录
	bucketSort := elastic.NewBucketSortAggregation().From(offset).Size(limit)

	// 将子聚合对象添加到主聚合中，命名为 "articles"
	agg.SubAggregation("articles", subAgg)
	// 将桶排序聚合对象添加到主聚合中，命名为 "page"
	agg.SubAggregation("page", bucketSort)

	result, err := global.ESClient.
		Search(esmodels.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		return nil, 0, err
	}

	var resp TagsAggregationResponse
	if err = json.Unmarshal(result.Aggregations["tags"], &resp); err != nil {
		return nil, 0, err
	}

	var aggregations []TagAggregation
	for _, bucket := range resp.Buckets {
		// 定义一个字符串切片，用于存储每个标签下的文章 ID 列表
		var articles []string
		for _, a := range bucket.Articles.Buckets {
			articles = append(articles, a.Key)
		}
		aggregations = append(aggregations, TagAggregation{
			Tag:           bucket.Key,
			DocCount:      bucket.DocCount,
			ArticleIDList: articles,
		})
	}
	return aggregations, total, nil
}

type T struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
	} `json:"buckets"`
}
