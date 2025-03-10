package fileService

import (
	"backend/global"
	"backend/models/diverseType"
	"backend/models/sqlmodels"
	"backend/plugins/freeimg"
	"backend/repository/img_repo"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xyproto/randomstring"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// FileUpload 单文件上传
func FileUpload(c *gin.Context, file *multipart.FileHeader, nickname string) FileUploadResponse {
	return handleUpload(c, file, nickname, diverseType.Local, localUploadHandler)
}

// FilesUpload 多文件上传
func FilesUpload(c *gin.Context, file *multipart.FileHeader, nickname string) FileUploadResponse {
	return handleUpload(c, file, nickname, diverseType.Local, localUploadHandler)
}

// FreeimgUpload 免费图床上传
func FreeimgUpload(c *gin.Context, file *multipart.FileHeader, nickname string) FileUploadResponse {
	res := handleUpload(c, file, nickname, diverseType.Remote, localUploadHandler)
	if !res.IsSuccess {
		return res
	}

	// 上传到图床并更新记录
	localPath := res.FileName
	remoteURL, err := freeimg.ImgCreate().Upload(localPath)
	if err != nil {
		_ = os.Remove(localPath)
		return errorResponse("图床上传失败,文件已存在: "+localPath, "")
	}

	if err = updateFileRecord(res.FileName, remoteURL, diverseType.Remote); err != nil {
		_ = os.Remove(localPath)
		return errorResponse("更新数据库失败", "")
	}

	_ = os.Remove(localPath)
	return successResponse(remoteURL, "图床上传成功"+"存储在:"+remoteURL)
}

// --- 内部辅助函数 ---
type uploadHandler func(*multipart.FileHeader, string, *gin.Context) error

func localUploadHandler(file *multipart.FileHeader, path string, c *gin.Context) error {
	return c.SaveUploadedFile(file, path)
}

// func freeimgUploadHandler(file *multipart.FileHeader, path string, c *gin.Context) error {
// 	return c.SaveUploadedFile(file, path)
// }

func handleUpload(c *gin.Context, file *multipart.FileHeader, nickname string, imgType diverseType.ImageType, handler uploadHandler) FileUploadResponse {

	// 文件验证
	if msg, err := validateFile(file); err != nil {
		return errorResponse(msg, "")
	}

	// 创建用户目录
	if err := createUserDir(nickname); err != nil {
		return errorResponse("目录创建失败", "")
	}

	// 生成唯一文件名
	fileName := generateUniqueName(file)
	filePath := path.Join(global.Config.Upload.Path, nickname, fileName)

	// 检查文件哈希
	hash := calculateFileHash(file)
	if hash == "" {
		return errorResponse("文件哈希计算失败", "")
	}

	// 检查重复文件
	if exists, existing := checkExistingFile(hash); exists {
		return successResponse(existing.Path, "文件已存在")
	}

	// 执行上传操作
	if err := handler(file, filePath, c); err != nil {
		return errorResponse("文件保存失败", filePath)
	}

	// 数据库存储
	if err := saveFileRecord(hash, filePath, fileName, imgType); err != nil {
		_ = os.Remove(filePath)
		return errorResponse("数据库存储失败", filePath)
	}

	return successResponse(filePath, "上传成功"+"存储在:"+filePath)
}

func validateFile(file *multipart.FileHeader) (string, error) {
	// 文件类型检查
	ext := strings.ToLower(path.Ext(file.Filename))
	if ext == "" || !utils.InList(ext[1:], WhiteImageList) {
		return "不支持的文件类型", ErrInvalidFile
	}

	// 文件大小检查
	if sizeMB := float64(file.Size) / (1024 * 1024); sizeMB >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("文件大小超标 (%.2fMB/%.dMB)", sizeMB, global.Config.Upload.Size)
		return msg, ErrInvalidFile
	}
	return "", nil
}

func createUserDir(nickname string) error {
	dirPath := path.Join(global.Config.Upload.Path, nickname)
	return os.MkdirAll(dirPath, 0755)
}

func generateUniqueName(file *multipart.FileHeader) string {
	ext := path.Ext(file.Filename)
	base := strings.TrimSuffix(file.Filename, ext)
	return fmt.Sprintf("%s_%s%s", base, randomstring.HumanFriendlyString(4), ext)
}

func calculateFileHash(file *multipart.FileHeader) string {
	f, err := file.Open()
	if err != nil {
		global.Log.Error("文件打开失败:", err)
		return ""
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		global.Log.Error("文件读取失败:", err)
		return ""
	}
	return utils.Md5(data)
}

func checkExistingFile(hash string) (bool, sqlmodels.BannerModel) {
	var banner sqlmodels.BannerModel
	if err := banner.GetByHash(hash); err != nil {
		return false, banner
	}
	global.Log.Info(banner)
	return true, banner
}

func saveFileRecord(hash, path, name string, imgType diverseType.ImageType) error {
	return img_repo.CreateBanner(sqlmodels.BannerModel{
		Path:      path,
		Hash:      hash,
		Name:      name,
		ImageType: imgType,
	})
}

func updateFileRecord(oldPath, newPath string, imgType diverseType.ImageType) error {
	var banner sqlmodels.BannerModel
	if err := banner.GetByPath(oldPath); err != nil {
		return err
	}
	return banner.UpdateBanner(map[string]any{
		"Path":      newPath,
		"ImageType": imgType,
	})
}

func errorResponse(msg, fileName string) FileUploadResponse {
	return FileUploadResponse{
		Msg:       msg,
		FileName:  fileName,
		IsSuccess: false,
	}
}

func successResponse(fileName, msg string) FileUploadResponse {
	return FileUploadResponse{
		Msg:       msg,
		FileName:  fileName,
		IsSuccess: true,
	}
}
