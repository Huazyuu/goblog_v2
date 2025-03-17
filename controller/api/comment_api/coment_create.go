package comment_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/service/commentService"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr req.CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := commentService.CommentCreateService(claims.UserID, cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
