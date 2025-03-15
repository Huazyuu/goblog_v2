package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/service/articleService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCollListView(c *gin.Context) {
	var cr req.PageInfo

	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	list, msg, err := articleService.ArticleCollListService(cr, claims.UserID)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithList(list, int64(len(list)), c)
}
