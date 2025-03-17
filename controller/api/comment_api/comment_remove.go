package comment_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/service/commentService"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr req.CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	msg, err := commentService.CommentRemoveService(claims, cr.ID)
	if err != nil {
		res.FailWithMessage(msg, c)
		global.Log.Error(err.Error())
		return
	}
	res.OkWithMessage(msg, c)
}
