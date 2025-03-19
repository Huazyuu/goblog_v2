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

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	// 判断是否本人 es user_id -> user id
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	title := "更新文章"
	log := logStash.NewAction(c)

	var cr req.ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := articleService.ArticleUpdateService(cr, claims.UserID)
	if err != nil {
		log.ErrItem(title, "更新文章错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "更新文章成功", fmt.Sprintf("用户%s更新文章%s", claims.Username, cr.Title))
	res.OkWithMessage(msg, c)
}
