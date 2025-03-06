package fileService

import (
	"backend/global"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"backend/plugins/freeimg"
	"backend/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func FileUpload(file *multipart.FileHeader, nickname string, c *gin.Context) (string, error) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		return "非法图片文件", errors.New("图片格式错误")
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		return msg, errors.New(msg)
	}
	dirList, err := os.ReadDir(basePath)
	if err != nil {
		return "文件目录不存在", errors.New("文件目录不存在")
	}
	if !isInDirEntry(dirList, nickname) {
		// 创建这个目录
		err := os.Mkdir(path.Join(basePath, nickname), 077)
		if err != nil {
			global.Log.Error(err)
		}
	}
	now := time.Now().Format("20060102150405")
	fileName = nameList[0] + "_" + now + "." + suffix
	filePath := path.Join(basePath, nickname, fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		return "保存失败", err
	}
	return fmt.Sprintf("图片保存在:%s", filePath), nil
}

func FilesUpload(file *multipart.FileHeader, nickname string) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, nickname, fileName)
	res.FileName = filePath
	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		return
	}

	// 读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)

	// 去数据库中查这个图片是否存在
	var bannerModel sqlmodels.BannerModel
	bannerModel.Hash = imageHash
	err = bannerModel.IsBannerExistByHash()
	if err == nil {
		// 找到了
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	bannerModel.Path = filePath
	bannerModel.Name = fileName
	bannerModel.ImageType = diverseType.Local

	err = bannerModel.CreateBanner()
	if err != nil {
		return FileUploadResponse{}
	}
	// 图片入库
	return res
}

func FreeimgUpload(c *gin.Context, file *multipart.FileHeader, nickname string) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, nickname, fileName)
	res.FileName = filePath
	// 读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)

	// 去数据库中查这个图片是否存在
	var bannerModel sqlmodels.BannerModel
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	err = c.SaveUploadedFile(file, res.FileName)
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		res.IsSuccess = false
	}
	krUpload := freeimg.ImgCreate().Upload(res.FileName)
	os.Remove(res.FileName)

	bannerModel.Hash = imageHash
	bannerModel.Path = krUpload
	bannerModel.Name = fileName
	bannerModel.ImageType = 3

	err = bannerModel.CreateBanner()
	if err != nil {
		return FileUploadResponse{}
	}
	res.FileName = krUpload
	// 图片入库
	return res
}

// isInDirEntry 检查指定名称的目录是否存在于给定的 os.DirEntry 切片中
func isInDirEntry(dirList []os.DirEntry, name string) bool {
	for _, entry := range dirList {
		if entry.Name() == name && entry.IsDir() {
			return true
		}
	}
	return false
}
