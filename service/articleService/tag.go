package articleService

import (
	"backend/controller/resp"
	"backend/repository/article_repo"
	"backend/repository/tag_repo"
)

func GetArticleTags(page, limit int) ([]resp.ArticleTagResponse, int64, error) {
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
	var response []resp.ArticleTagResponse
	for _, agg := range aggregations {
		response = append(response, resp.ArticleTagResponse{
			Tag:           agg.Tag,
			Count:         agg.DocCount,
			ArticleIDList: agg.ArticleIDList,
			CreatedAt:     createTimes[agg.Tag],
		})
	}

	return response, total, nil
}
