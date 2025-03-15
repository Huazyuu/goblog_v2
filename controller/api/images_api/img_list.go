package images_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/service/fileService"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr req.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	imageList, err := fileService.ImageListService(cr)
	res.OkWithList(imageList, int64(len(imageList)), c)
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
	res.OkWithList(imageList, int64(len(imageList)), c)
}
