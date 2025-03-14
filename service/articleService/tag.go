package articleService

import (
	"backend/repository/article_repo"
	"backend/repository/tag_repo"
)

type TagResponse struct {
	Tag           string   `json:"tag"`
	Count         int64    `json:"count"`
	ArticleIDList []string `json:"articleIDList"`
	CreatedAt     string   `json:"createdAt"`
}

func GetArticleTags(page, limit int) ([]TagResponse, int64, error) {
	// 获取ES聚合数据
	aggregations, total, err := article_repo.GetTagAggregations(page, limit)
	if err != nil {
		return nil, 0, err
	}
	// 收集标签列表
	var tags []string
	for _, agg := range aggregations {
		tags = append(tags, agg.Tag)
	}
	// 获取标签创建时间
	createTimes, err := tag_repo.GetTagCreateTimes(tags)
	if err != nil {
		return nil, 0, err
	}
	// 组装响应数据
	var response []TagResponse
	for _, agg := range aggregations {
		response = append(response, TagResponse{
			Tag:           agg.Tag,
			Count:         agg.DocCount,
			ArticleIDList: agg.ArticleIDList,
			CreatedAt:     createTimes[agg.Tag],
		})
	}

	return response, total, nil
}
