package article_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/repository/article_repo"
	"backend/service/redisService"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDiggView(c *gin.Context) {
	var cr req.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	exist, err := article_repo.ISArticleExistByID(cr.ID)
	if !exist {
		global.Log.Error(err)
		res.FailWithMessage("没有该文章", c)
		return
	}
	// 查es
	err = redisService.NewArticleDigg().Set(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("点赞出错 err: "+err.Error(), c)
		return
	}
	res.OkWithMessage("文章点赞成功", c)
}
