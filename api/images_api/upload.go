package images_api

import (
	"backend/global"
	"backend/middleware/jwt"
	"backend/models/res"
	"backend/service/imgService"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

// ImageUploadView 上传单个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("参数校验失败", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.Role == 3 {
		res.FailWithMessage("游客用户不可上传图片", c)
		return
	}

	msg, err := fileService.FileUpload(file, claims.NickName, c)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(msg, c)

}

// ImagesUploadView 上传多个图片，返回图片的url
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.Role == 3 {
		res.FailWithMessage("游客用户不可上传图片", c)
		return
	}
	// 上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	// 判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []fileService.FileUploadResponse

	for _, file := range fileList {
		// 上传文件
		serviceRes := fileService.FilesUpload(file, claims.NickName)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 成功的
		if !global.Config.QiNiu.Enable {
			// 本地还得保存一下
			global.Log.Info(serviceRes.FileName)
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}

func (ImagesApi) FreeImagesUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.Role == 3 {
		res.FailWithMessage("游客用户不可上传图片", c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		res.FailWithMessage("读取图片失败", c)
		return
	}

	serviceRes := fileService.FreeimgUpload(c, file, claims.NickName)

	res.OkWithData(serviceRes, c)
}
