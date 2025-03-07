package adverts_api

import (
	"backend/models/common"
	"backend/models/res"
	"backend/models/sqlmodels"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AdvertsApi) AdvertListView(c *gin.Context) {
	var cr common.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	referer := c.GetHeader("Gvb_referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		// admin来的
		isShow = false
	}
	// todo [bug] should by fix , when isShow is false ,gorm select will ignore this field (false is 0,the null value)
	list, count, _ := common.ComList(sqlmodels.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
