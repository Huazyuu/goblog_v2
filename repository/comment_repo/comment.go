package comment_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
	"backend/service/redisService"
	"fmt"
	"gorm.io/gorm"
)

func GetCommentByID(id uint) (sqlmodels.CommentModel, error) {
	var model sqlmodels.CommentModel
	err := global.DB.Where("id = ?", id).Take(&model).Error
	return model, err
}
func UpdateComment(id uint, mapdata map[string]interface{}) error {
	var comment sqlmodels.CommentModel
	result := global.DB.Model(&comment).Where("id = ?", id).Updates(mapdata)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func CreateComment(user sqlmodels.CommentModel) error {
	return global.DB.Create(&user).Error
}

func GetCountGroupByArticle() (int64, error) {
	var count int64
	err := global.DB.Model(&sqlmodels.CommentModel{}).Group("article_id").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

type ArticleIdCount struct {
	ArticleID string
	Count     int
}

func GetArticleIdCount(page, limit int) ([]ArticleIdCount, error) {
	global.Log.Info(page)
	global.Log.Info(limit)
	offset := (page - 1) * limit
	var list []ArticleIdCount
	res := global.DB.Model(sqlmodels.CommentModel{}).
		Group("article_id").
		Order("count desc").
		Limit(limit).
		Offset(offset).
		Select("article_id", "count(id) as count").
		Scan(&list)
	err := res.Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// GetArticleCommentList 获取指定文章的评论列表，包含根评论及其子评论树
func GetArticleCommentList(articleID string) (RootCommentList []*sqlmodels.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)

	diggInfo := redisService.NewCommentDigg().GetAll()

	// 遍历根评论，递归查根评论下的所有子评论，并更新点赞数
	for _, model := range RootCommentList {
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		model.DiggCount = model.DiggCount + modelDigg
		getCommentTree(model)
	}
	return
}

// GetAllSubCommentList 查找一个评论的所有子评论，并将其一维化
func GetAllSubCommentList(comment sqlmodels.CommentModel) (subList []sqlmodels.CommentModel) {
	global.DB.Preload("SubComments").Preload("User").Take(&comment)

	for _, model := range comment.SubComments {
		// 将当前子评论添加到结果切片中
		subList = append(subList, *model)
		// 递归查找当前子评论的所有子评论，并添加到结果切片中
		subList = append(subList, GetAllSubCommentList(*model)...)
	}
	return
}

// getCommentTree 获取指定评论的评论树
func getCommentTree(rootComment *sqlmodels.CommentModel) *sqlmodels.CommentModel {
	// 使用全局数据库连接，预加载评论的用户信息和子评论信息，获取指定评论
	global.DB.Preload("User").Preload("SubComments").Find(rootComment)

	// 递归获取子评论树
	for _, subComment := range rootComment.SubComments {
		// 递归调用 getCommentTree 方法，获取当前子评论的评论树
		if subComment == nil {
			return nil
		}
		getCommentTree(subComment)
	}
	return rootComment
}

func IsExistCommentID(id uint) bool {
	var count int64
	global.DB.Model(&sqlmodels.CommentModel{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func DeleteComments(commentIDs []uint) (count int64, err error) {
	var comments []sqlmodels.CommentModel
	// 根据 ID 列表删除评论
	result := global.DB.Where("id IN ?", commentIDs).Delete(&comments)
	count = result.RowsAffected
	err = result.Error
	return count, err
}
