package adverts_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/models/sqlmodels"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AdvertsApi) AdvertListView(c *gin.Context) {
	var cr req.PageInfo
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
	list, count, _ := req.ComList(sqlmodels.AdvertModel{IsShow: isShow}, req.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}
