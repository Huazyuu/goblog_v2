package resp

import "time"

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
type CommentByArticleListResponse struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Count int    `json:"count"`
}
