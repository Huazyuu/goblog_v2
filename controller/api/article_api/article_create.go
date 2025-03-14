package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"

	"backend/service/articleService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr req.ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := articleService.ArticleService(cr, claims)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
