package flags

import (
	"backend/core"
	"backend/global"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// getProjectRoot 获取项目绝对根路径（包含go.mod的目录）
func getProjectRoot() string {
	// 获取当前测试文件路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("无法获取调用者信息")
	}

	// 向上递归查找包含 go.mod 的目录
	currentDir := filepath.Dir(filename)
	for {
		// 检查当前目录是否包含 go.mod
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir
		}

		// 到达文件系统根目录时终止
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			log.Fatal("已到达文件系统根目录，未找到go.mod")
		}
		currentDir = parentDir
	}
}
func init() {
	// 获取并切换到项目根目录
	root := getProjectRoot()
	if err := os.Chdir(root); err != nil {
		log.Fatalf("切换工作目录失败: %v", err)
	}

	// 初始化应用配置（现在可以使用相对路径）
	global.Config = core.InitConf() // 根目录下的配置文件
	global.Log = core.InitLogger()
	global.ESClient = core.InitES()
}

func Test_esCreate(t *testing.T) {
	esCreate()
}
