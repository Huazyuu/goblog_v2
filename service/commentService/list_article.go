package commentService

import (
	"backend/controller/req"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
	"backend/repository/comment_repo"
	"github.com/liu-cn/json-filter/filter"
)

type CommentByArticleListResponse struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Count int    `json:"count"`
}

func CommentByArticleListService(cr req.PageInfo) ([]CommentByArticleListResponse, error) {
	articleIDlist, err := comment_repo.GetArticleIdCount(cr.Page, cr.Limit)
	if err != nil {
		return nil, err
	}
	var articleIDMap = make(map[string]int)
	var articleIDList []any
	for _, a := range articleIDlist {
		articleIDMap[a.ArticleID] = a.Count
		articleIDList = append(articleIDList, a.ArticleID)
	}
	articles, err := article_repo.GetArticleListByIDList(articleIDList)
	if err != nil {
		return nil, err
	}
	list := make([]CommentByArticleListResponse, 0, len(articles))
	for _, article := range articles {
		list = append(list, CommentByArticleListResponse{
			Title: article.Title,
			ID:    article.ID,
			Count: articleIDMap[article.ID],
		})
	}
	return list, nil
}

func CommentListInArticle(articleID string) (any, int64) {
	CommentList := comment_repo.GetArticleCommentList(articleID)
	data := filter.Select("c", CommentList)
	filterList := data.(filter.Filter)
	if string(filterList.MustMarshalJSON()) == "{}" {
		list := make([]sqlmodels.CommentModel, 0)
		return list, 0
	}
	return data, int64(len(CommentList))
}
