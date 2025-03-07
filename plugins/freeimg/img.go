package freeimg

import (
	"backend/global"
)

type ImgUploadInterface interface {
	Upload(filename string) (string, error)
}

var serveMap = map[string]ImgUploadInterface{
	"freeimg": &FreeImageService{},
}

func ImgCreate() ImgUploadInterface {
	return serveMap[global.Config.PictureBed.Type]
}
