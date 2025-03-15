package articleService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
	"fmt"
)

func ArticleCollRemoveService(cr req.ESIDListRequest, userid uint) (string, error) {
	articleIDList, colls, err := getCollArticleIDListByUserID(cr.IDList, userid)
	if err != nil || len(articleIDList) == 0 {
		return "非法请求", err
	}
	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}
	articleList, err := article_repo.GetArticleListByIDList(idList)
	if err != nil {
		return "系统错误:" + err.Error(), err
	}
	for _, article := range articleList {
		err = article_repo.UpdateArticle(article.ID, map[string]interface{}{
			"collects_count": article.CollectsCount - 1,
		})
		if err != nil {
			return "系统错误:" + err.Error(), err
		}
	}
	global.DB.Delete(colls)
	return fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), nil
}
func getCollArticleIDListByUserID(articleIDList []string, userID uint) (articleIDs []string, colls []sqlmodels.CollectModel, err error) {
	err = global.DB.
		Find(&colls, "user_id = ? and article_id in ?", userID, articleIDList).
		Select("article_id").
		Scan(&articleIDs).Error
	return
}
