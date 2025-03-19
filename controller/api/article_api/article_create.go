package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"
	"backend/plugins/logStash"
	"backend/service/articleService"
	"fmt"
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

	log := logStash.NewAction(c)

	msg, err := articleService.ArticleCreateService(cr, claims)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(msg, c)
		return
	}

	log.SetItemInfo("创建文章", fmt.Sprintf("用户%s创建文章%s", claims.Username, cr.Title))

	res.OkWithMessage(msg, c)
}
