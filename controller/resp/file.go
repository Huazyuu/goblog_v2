package resp

import "backend/models/sqlmodels"

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}
type ImageListResponse struct {
	sqlmodels.BannerModel
	BannerCount  int `json:"bannerCount"`  // 关联banner的个数
	ArticleCount int `json:"articleCount"` // 关联文章的个数
}
