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

func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr req.ESIDRequest
	title := "收藏文章"
	log := logStash.NewAction(c)

	err := c.ShouldBindUri(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	msg, err := articleService.ArticleCollCreateService(cr.ID, claims.UserID)
	if err != nil {
		log.ErrItem(title, "收藏错误", err.Error())
		global.Log.Error(err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.InfoItem(title, "收藏文章成功", fmt.Sprintf("用户%s收藏文章id:%s", claims.Username, cr.ID))
	res.OkWithMessage(msg, c)
}
