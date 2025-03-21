package resp

import "backend/models/esmodels"

type ArticleCollResponse struct {
	esmodels.ArticleModel
	CreatedAt string `json:"created_at"`
}

type ArticleItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ArticleDetailResponse struct {
	esmodels.ArticleModel
	IsCollect bool         `json:"is_collect"` // 用户是否收藏文章
	Next      *ArticleItem `json:"next"`       // 上一篇文章
	Prev      *ArticleItem `json:"prev"`       // 下一篇文章
}
type ArticleCategoryResponse struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}
type ArticleTagResponse struct {
	Tag           string   `json:"tag"`
	Count         int64    `json:"count"`
	ArticleIDList []string `json:"articleIDList"`
	CreatedAt     string   `json:"createdAt"`
}
