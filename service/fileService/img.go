package fileService

import "errors"

var (
	WhiteImageList = []string{"jpg", "png", "jpeg", "ico", "tiff", "gif", "svg", "webp"}
	ErrInvalidFile = errors.New("invalid file")
)
