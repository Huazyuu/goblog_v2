package images_api

import (
	"backend/global"
	"backend/models/common"
	"backend/models/res"
	"backend/models/sqlmodels"
	"github.com/gin-gonic/gin"
)

// todo ImageListView 图片列表 文章 菜单图片关联

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr common.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	imageList, cnt, err := common.ComList(sqlmodels.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Likes:    []string{"name", "path"},
	})
	res.OkWithList(imageList, cnt, c)
}

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` // 图片路径
	Name string `json:"name"` // 图片名称
}

// ImageNameListView 图片名称列表
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(sqlmodels.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)
}
