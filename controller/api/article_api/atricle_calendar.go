package article_api

import (
	"backend/controller/res"
	"backend/global"
	"backend/models/esmodels"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"time"
)

type calendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type bucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var dateCount = make(map[string]int)

func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	now := time.Now()
	past := now.AddDate(-1, 0, 0)
	format := "2006-01-02 15:04:05"

	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	query := elastic.NewRangeQuery("created_at").Gte(past.Format(format)).Lte(now.Format(format))
	result, err := global.ESClient.Search(esmodels.ArticleModel{}.Index()).Query(query).Aggregation("bucket", agg).Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var data bucketsType
	_ = json.Unmarshal(result.Aggregations["bucket"], &data)

	reslist := make([]calendarResponse, 0)
	for _, bucket := range data.Buckets {
		t, _ := time.Parse(format, bucket.KeyAsString)
		dateCount[t.Format("2006-01-02")] = bucket.DocCount
	}
	days := int(now.Sub(past).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := past.AddDate(0, 0, i).Format("2006-01-02")
		count, _ := dateCount[day]
		if count != 0 {
			reslist = append(reslist, calendarResponse{
				Date:  day,
				Count: count,
			})
		}
	}
	res.OkWithData(reslist, c)
}
