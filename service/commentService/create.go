package commentService

import (
	"backend/controller/req"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
	"backend/repository/comment_repo"
	"backend/service/redisService"
	"errors"
	"gorm.io/gorm"
)

func CommentCreateService(uid uint, cr req.CommentRequest) (string, error) {
	ok, err := article_repo.IsExistArticleByID(cr.ArticleID)
	if err != nil {
		return "查找错误", err
	}
	if !ok {
		return "没有这篇文章", errors.New("没有这篇文章")
	}

	if cr.ParentCommentID != nil {
		// 是子评论
		parent, err := comment_repo.GetCommentByID(*cr.ParentCommentID)
		if err != nil {
			return "父评论不存在", err
		}
		// 判断父评论的文章是否和当前文章一致
		if parent.ArticleID != cr.ArticleID {
			return "评论文章不一致", err
		}
		// 父 cnt++
		err = comment_repo.UpdateComment(*cr.ParentCommentID, map[string]interface{}{
			"comment_count": gorm.Expr("comment_count + 1"),
		})
		if err != nil {
			return "更新错误", err
		}
	}

	// 不是子评论
	err = comment_repo.CreateComment(sqlmodels.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          uid,
	})
	if err != nil {
		return "创建失败", err
	}
	redisService.NewCommentCount().Get(cr.ArticleID)
	return "创建成功", nil
}
