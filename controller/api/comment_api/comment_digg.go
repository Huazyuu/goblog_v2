package comment_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/plugins/logStash"
	"backend/repository/comment_repo"
	"backend/service/redisService"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (CommentApi) CommentDiggView(c *gin.Context) {
	var cr req.CommentIDRequest
	title := "点赞评论"
	log := logStash.NewAction(c)

	err := c.ShouldBindUri(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if ok := comment_repo.IsExistCommentID(cr.ID); !ok {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage("评论不存在", c)
		return
	}
	redisService.NewCommentDigg().Set(strconv.Itoa(int(cr.ID)))
	log.InfoItem(title, "点赞评论成功", fmt.Sprintf("点赞评论id:%d", cr.ID))
	res.OkWithMessage("评论点赞成功", c)
	return
}
