package freeimg

import (
	"backend/core"
	"backend/global"
	"testing"
)

func Test_freeImgUpload(t *testing.T) {
	global.Config = core.InitConf("../../settings.yaml")
	global.Log = core.InitLogger()
	// 替换为实际的本地文件路径
	uploadFile := "D:\\goProject\\goblogv2\\backend\\uploads\\avatar\\default.png"
	url := freeImgUpload(uploadFile)
	if url != "" {
		t.Logf("上传成功，图片显示 URL: %s\n", url)
	} else {
		t.Fatal("上传失败")
		return
	}
}
