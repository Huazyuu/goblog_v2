package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/middleware/jwt"
	"backend/plugins/logStash"
	"backend/service/articleService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr req.ArticleRequest
	log := logStash.NewAction(c)
	title := "创建文章"
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := articleService.ArticleCreateService(cr, claims)
	if err != nil {
		log.ErrItem(title, "创建出错", err.Error())
		res.FailWithMessage(msg, c)
		return
	}

	log.InfoItem(title, "创建文章成功", fmt.Sprintf("用户%s创建文章%s", claims.Username, cr.Title))

	res.OkWithMessage(msg, c)
}
