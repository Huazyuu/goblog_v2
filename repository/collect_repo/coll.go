package collect_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

func GetCollByUserIDAndArticleID(articleID string, userID uint) (coll sqlmodels.CollectModel, err error) {
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", userID, articleID).Error
	return
}
func GetCollListByUserID(articleIDList []string, userID uint) (colls []sqlmodels.CollectModel, err error) {
	err = global.DB.Find(&colls, "user_id = ? and article_id in ?", userID, articleIDList).Error
	return
}
func GetCollArticleIDListByUserID(articleIDList []string, userID uint) (articleIDs []string, colls []sqlmodels.CollectModel, err error) {
	err = global.DB.
		Find(&colls, "user_id = ? and article_id in ?", userID, articleIDList).
		Select("article_id").
		Scan(&articleIDs).Error
	return
}

func CollectArticle(articleID string, userID uint) error {
	return global.DB.Create(&sqlmodels.CollectModel{
		UserID:    userID,
		ArticleID: articleID,
	}).Error
}
func RemoveCollect(coll *sqlmodels.CollectModel) error {
	return global.DB.Delete(coll).Error
}
func RemoveCollectsByID(coll *sqlmodels.CollectModel) error {
	return global.DB.Delete(coll).Error
}
