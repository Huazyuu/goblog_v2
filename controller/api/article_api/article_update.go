package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/service/articleService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	// 判断是否本人 es user_id -> user id
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr req.ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := articleService.ArticleUpdateService(cr, claims.UserID)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
