package commentService

import (
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/sqlmodels"
	"backend/repository/comment_repo"
	"backend/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func CommentRemoveService(claims *jwt.CustomClaims, commentid uint) (string, error) {
	comment, err := comment_repo.GetCommentByID(commentid)
	if err != nil {
		return "评论不存在", err
	}
	// 这条评论只能由当前登录人删除，或者管理员
	if !(comment.UserID == claims.UserID || claims.Role == 1) {
		return "权限错误，不可删除", errors.New("权限错误，不可删除")
	}
	// 获得子评论
	subCommentList := comment_repo.GetAllSubCommentList(comment)
	// 评论总数 包括自己
	cnt := len(subCommentList) + 1
	// 是子评论
	if comment.ParentCommentID != nil {
		err = comment_repo.UpdateComment(*comment.ParentCommentID, map[string]interface{}{
			"comment_count": gorm.Expr("comment_count - ?", cnt),
		})
	}
	// 不是,删除子评论以及当前评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, comment.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(sqlmodels.CommentModel{}).Delete("id = ?", id)
	}
	return fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), nil
}
