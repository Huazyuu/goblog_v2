package commentService

import (
	"backend/controller/req"
	"backend/models/esmodels"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
	"time"
)

type CommentListResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"` // 创建时间
	ArticleTitle    string    `json:"article_title"`
	ArticleBanner   string    `json:"article_banner"`
	ParentCommentID *uint     `json:"parent_comment_id"`
	Content         string    `json:"content"`
	DiggCount       int       `json:"digg_count"`
	CommentCount    int       `json:"comment_count"`
	UserNickName    string    `json:"user_nick_name"`
}

func CommentListService(cr req.PageInfo) ([]CommentListResponse, error) {
	list, _, err := req.ComList(sqlmodels.CommentModel{}, req.Option{
		PageInfo: cr,
		Preload:  []string{"User"},
	})

	commentList := make([]CommentListResponse, 0, len(list))
	articleIDList := make([]any, 0, len(list))
	articleMap := make(map[string]esmodels.ArticleModel)

	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
	}

	articles, err := article_repo.GetArticleListByIDList(articleIDList)
	if err != nil {
		return nil, err
	}

	for _, article := range articles {
		articleMap[article.ID] = article
	}

	for _, model := range list {
		commentList = append(commentList, CommentListResponse{
			ID:              model.ID,
			CreatedAt:       model.CreatedAt,
			ParentCommentID: model.ParentCommentID,
			Content:         model.Content,
			DiggCount:       model.DiggCount,
			CommentCount:    model.CommentCount,
			UserNickName:    model.User.NickName,
			ArticleTitle:    articleMap[model.ArticleID].Title,
			ArticleBanner:   articleMap[model.ArticleID].BannerUrl,
		})
	}
	return commentList, nil

}
