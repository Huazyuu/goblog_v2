package images_api

import (
	"backend/controller/res"
	"backend/global"
	"backend/middleware/jwt"

	"backend/service/fileService"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// ImageUploadView 单文件上传
func (ImagesApi) ImageUploadView(c *gin.Context) {
	handleUpload(c, func(c *gin.Context, claims *jwt.CustomClaims) interface{} {
		return singleFileUpload(c, claims)
	})
}

// ImagesUploadView 多文件上传
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	handleUpload(c, func(c *gin.Context, claims *jwt.CustomClaims) interface{} {
		return multipleFilesUpload(c, claims)
	})
}

// FreeImagesUploadView 免费图床上传
func (ImagesApi) FreeImagesUploadView(c *gin.Context) {
	handleUpload(c, func(c *gin.Context, claims *jwt.CustomClaims) interface{} {
		return freeImageUpload(c, claims)
	})
}

// --- 核心处理逻辑 ---

// 统一上传入口
func handleUpload(c *gin.Context, handler func(*gin.Context, *jwt.CustomClaims) interface{}) {
	// 统一权限验证
	claims, ok := validateUserRole(c)
	if !ok {
		return
	}

	// 执行具体上传逻辑
	result := handler(c, claims)

	// 统一响应处理
	switch v := result.(type) {
	case fileService.FileUploadResponse:
		if v.IsSuccess {
			res.OkWithData(v, c)
		} else {
			res.FailWithMessage(v.Msg, c)
		}
	case []fileService.FileUploadResponse:
		res.OkWithData(v, c)
	case error:
		res.FailWithMessage(v.Error(), c)
	default:
		res.OkWithData(v, c)
	}
}

// 验证用户权限
func validateUserRole(c *gin.Context) (*jwt.CustomClaims, bool) {
	claims, err := getJWTClaims(c)
	if err != nil {
		res.FailWithMessage("用户信息获取失败", c)
		return nil, false
	}

	if claims.Role == 3 {
		res.FailWithMessage("游客用户不可上传图片", c)
		return nil, false
	}
	return claims, true
}

// --- 具体上传类型处理 ---

// 单文件上传处理
func singleFileUpload(c *gin.Context, claims *jwt.CustomClaims) interface{} {
	file, err := getSingleFile(c)
	if err != nil {
		return err
	}
	return fileService.FileUpload(c, file, claims.NickName)
}

// 多文件上传处理
func multipleFilesUpload(c *gin.Context, claims *jwt.CustomClaims) interface{} {
	files, err := getMultipleFiles(c)
	if err != nil {
		return err
	}

	var results []fileService.FileUploadResponse
	for _, file := range files {
		result := fileService.FilesUpload(c, file, claims.NickName)
		results = append(results, result)
	}
	return results
}

// 免费图床上传处理
func freeImageUpload(c *gin.Context, claims *jwt.CustomClaims) interface{} {
	file, err := getSingleFile(c)
	if err != nil {
		return err
	}
	return fileService.FreeimgUpload(c, file, claims.NickName)
}

// --- 工具函数 ---

func getJWTClaims(c *gin.Context) (*jwt.CustomClaims, error) {
	_claims, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("用户信息不存在")
	}
	return _claims.(*jwt.CustomClaims), nil
}

func getSingleFile(c *gin.Context) (*multipart.FileHeader, error) {
	file, err := c.FormFile("file")
	if err != nil {
		global.Log.Error("文件获取失败:", err)
		return nil, errors.New("文件参数错误")
	}
	return file, nil
}

func getMultipleFiles(c *gin.Context) ([]*multipart.FileHeader, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, errors.New("表单解析失败")
	}

	files, ok := form.File["images"]
	if !ok || len(files) == 0 {
		return nil, errors.New("未找到有效文件")
	}
	return files, nil
}
