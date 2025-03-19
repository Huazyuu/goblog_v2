package comment_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/plugins/logStash"
	"backend/service/commentService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr req.CommentIDRequest
	err := c.ShouldBindUri(&cr)
	title := "删除评论"
	log := logStash.NewAction(c)

	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	msg, err := commentService.CommentRemoveService(claims, cr.ID)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		res.FailWithMessage(msg, c)
		global.Log.Error(err.Error())
		return
	}
	log.WarnItem(title, "删除评论成功", fmt.Sprintf("用户%s删除评论id:%d", claims.Username, cr.ID))
	res.OkWithMessage(msg, c)
}
