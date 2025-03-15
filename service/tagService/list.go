package tagService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/esmodels"
	"backend/models/sqlmodels"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"strconv"
)

func TagList(cr req.PageInfo) ([]sqlmodels.TagModel, int64, error) {
	list, cnt, err := req.ComList(sqlmodels.TagModel{}, req.Option{
		PageInfo: cr,
	})
	if err != nil {
		return nil, 0, err
	}
	return list, cnt, nil
}

type TagResponse struct {
	TagName string `json:"tag_name"`
	Count   string `json:"count"`
}

func TagNameListService() ([]TagResponse, error) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	query := elastic.NewBoolQuery()
	agg := elastic.NewTermsAggregation().Field("tags")
	result, err := global.ESClient.
		Search(esmodels.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	byteData := result.Aggregations["tags"]
	var tagType T
	json.Unmarshal(byteData, &tagType)

	var tagList = make([]TagResponse, 0)
	for _, bucket := range tagType.Buckets {
		tagList = append(tagList, TagResponse{
			TagName: bucket.Key,
			Count:   strconv.Itoa(bucket.DocCount),
		})
	}
	return tagList, nil
}
