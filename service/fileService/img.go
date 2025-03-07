package fileService

import "errors"

var (
	WhiteImageList = []string{"jpg", "png", "jpeg", "ico", "tiff", "gif", "svg", "webp"}
	ErrInvalidFile = errors.New("invalid file")
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}
