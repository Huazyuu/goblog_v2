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

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr req.CommentRequest
	title := "发布评论"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := commentService.CommentCreateService(claims.UserID, cr)
	if err != nil {
		log.ErrItem(title, "服务错误", err.Error())
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	log.InfoItem(title, "创建评论成功", fmt.Sprintf("用户%s发表评论%s", claims.Username, cr.Content))
	res.OkWithMessage(msg, c)
}
