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

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr req.ESIDListRequest
	title := "删除文章"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	msg, err := articleService.ArticleRemoveService(cr, claims)
	if err != nil {
		global.Log.Error(err)
		log.ErrItem(title, "删除文章错误", err.Error())
		res.FailWithMessage(msg, c)
		return
	}
	log.WarnItem(title, "创建文章成功", fmt.Sprintf("用户%s创建文章ID列表:%s", claims.Username, cr.IDList))
	res.OkWithMessage(msg, c)
}
