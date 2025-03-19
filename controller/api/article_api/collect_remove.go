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

func (ArticleApi) ArticleCollRemoveView(c *gin.Context) {
	var cr req.ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	title := "取消收藏文章"
	log := logStash.NewAction(c)

	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := articleService.ArticleCollRemoveService(cr, claims.UserID)
	if err != nil {
		log.ErrItem(title, "取消收藏失败", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "取消收藏失败", fmt.Sprintf("用户%s取消收藏文章id列表:%s", claims.Username, cr.IDList))
	res.OkWithMessage(msg, c)
}
