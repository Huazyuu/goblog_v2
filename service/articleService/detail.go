package articleService

import (
	"backend/controller/req"
	"backend/middleware/jwt"
	"backend/models/esmodels"
	"backend/repository/article_repo"
	"backend/repository/collect_repo"
	"backend/service/redisService"
	"github.com/gin-gonic/gin"
	"strings"
)

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

func ArticleDetailService(c *gin.Context, articleID string) (ArticleDetailResponse, string, error) {
	redisService.NewArticleLook().Set(articleID)

	article, err := article_repo.GetArticleByID(articleID)
	if err != nil {
		return ArticleDetailResponse{}, "请输入正确文章id", err
	}

	isCollect := isUserArticleColl(c, article.ID)

	var articleDetail = ArticleDetailResponse{
		ArticleModel: article,
		IsCollect:    isCollect,
	}

	// 根据分类，查文章列表，然后找当前文章所在的位置
	list, _, err := article_repo.GetArticleList(article_repo.Option{
		PageInfo: req.PageInfo{
			Limit: 10000,
			Page:  1,
		},
		Category: article.Category,
	})
	if err != nil {
		return ArticleDetailResponse{}, "查找错误", err
	}
	var currentIndex = -1
	// 查找当前记录的索引
	for i, item := range list {
		if item.ID == article.ID {
			currentIndex = i
			break
		}
	}

	var previousArticle esmodels.ArticleModel
	var nextArticle esmodels.ArticleModel

	if currentIndex > 0 {
		previousArticle = list[currentIndex-1]
		articleDetail.Next = &ArticleItem{
			ID:    previousArticle.ID,
			Title: previousArticle.Title,
		}
	}

	// 查找下一个记录
	if currentIndex < len(list)-1 {
		nextArticle = list[currentIndex+1]
		articleDetail.Prev = &ArticleItem{
			ID:    nextArticle.ID,
			Title: nextArticle.Title,
		}
	}
	return articleDetail, "查找成功", nil
}
func isUserArticleColl(c *gin.Context, articleID string) (isCollect bool) {
	// 判断用户是否正常登录
	authHeader := c.GetHeader("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return
	}
	tokenString := parts[1]
	if tokenString == "" {
		return
	}
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		return
	}
	// 判断是否在redis中
	if redisService.CheckLogout(tokenString) {
		return
	}
	_, err = collect_repo.GetCollByUserIDAndArticleID(articleID, claims.UserID)
	if err != nil {
		return
	}
	return true
}
