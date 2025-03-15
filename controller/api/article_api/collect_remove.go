package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/middleware/jwt"
	"backend/service/articleService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCollRemoveView(c *gin.Context) {
	var cr req.ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := articleService.ArticleCollRemoveService(cr, claims.UserID)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
