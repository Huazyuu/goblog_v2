package articleService

import (
	"backend/repository/article_repo"
	"backend/repository/collect_repo"
)

func ArticleCollCreateService(articleID string, userID uint) (string, error) {
	article, err := article_repo.GetArticleByID(articleID)
	if err != nil {
		return "文章不存在", err
	}

	coll, err := collect_repo.GetCollByUserIDAndArticleID(articleID, userID)
	num := -1
	if err != nil {
		// 没有找到 收藏文章
		err = collect_repo.CollectArticle(articleID, userID)
		if err != nil {
			return "收藏失败", err
		}
		// 给文章的收藏数 +1
		num = 1
	} else {
		// 取消收藏
		err = collect_repo.RemoveCollect(&coll)
		if err != nil {
			return "系统错误", err
		}
	}

	err = article_repo.UpdateArticle(articleID, map[string]any{
		"collects_count": article.CollectsCount + num,
	})
	if num == 1 {
		return "收藏成功", nil
	} else {
		return "取消成功", nil
	}
}
