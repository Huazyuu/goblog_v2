package articleService

import (
	"backend/controller/req"
	"backend/models/esmodels"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
)

type CollResponse struct {
	esmodels.ArticleModel
	CreatedAt string `json:"created_at"`
}

func ArticleCollListService(cr req.PageInfo, userid uint) ([]CollResponse, string, error) {
	var articleIDList []any
	var collMap = make(map[string]string)
	var collList = make([]CollResponse, 0)

	list, _, err := req.ComList(sqlmodels.CollectModel{UserID: userid}, req.Option{
		PageInfo: cr,
	})
	for _, coll := range list {
		articleIDList = append(articleIDList, coll.ArticleID)
		collMap[coll.ArticleID] = coll.CreatedAt.Format("2006-01-02 15:04:05")
	}
	articleList, err := article_repo.GetArticleListByIDList(articleIDList)
	if err != nil {
		return nil, "系统错误:" + err.Error(), err
	}
	for _, article := range articleList {
		collList = append(collList, CollResponse{
			ArticleModel: article,
			CreatedAt:    collMap[article.ID],
		})
	}
	return collList, "查询成功", nil

}
