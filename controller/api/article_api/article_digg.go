package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/plugins/logStash"
	"backend/repository/article_repo"
	"backend/service/redisService"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDiggView(c *gin.Context) {
	var cr req.ESIDRequest
	title := "文章点赞"
	log := logStash.NewAction(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.ErrItem(title, "参数绑定错误", err.Error())
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	exist, err := article_repo.IsExistArticleByID(cr.ID)
	if !exist {
		global.Log.Error(err)
		log.ErrItem(title, "没有该文章", nil)
		res.FailWithMessage("没有该文章", c)
		return
	}
	// 查es
	err = redisService.NewArticleDigg().Set(cr.ID)
	if err != nil {
		global.Log.Error(err)
		log.ErrItem(title, "点赞出错", err)
		res.FailWithMessage("点赞出错 err: "+err.Error(), c)
		return
	}
	log.InfoItem(title, "文章点赞成功", fmt.Sprintf("点赞文章id:%s", cr.ID))
	res.OkWithMessage("文章点赞成功", c)
}
