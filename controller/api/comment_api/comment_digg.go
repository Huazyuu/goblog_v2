package comment_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/repository/comment_repo"
	"backend/service/redisService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (CommentApi) CommentDiggView(c *gin.Context) {
	var cr req.CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if ok := comment_repo.IsExistCommentID(cr.ID); !ok {
		res.FailWithMessage("评论不存在", c)
		return
	}
	redisService.NewCommentDigg().Set(strconv.Itoa(int(cr.ID)))
	res.OkWithMessage("评论点赞成功", c)
	return
}
