package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/service/articleService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr req.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	msg, err := articleService.ArticleCollCreateService(cr.ID, claims.UserID)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
