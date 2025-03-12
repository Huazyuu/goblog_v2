package esmodels

import (
	"backend/core"
	"backend/global"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

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
func TestArticleModel(t *testing.T) {
	model := ArticleModel{}

	// 清理环境，确保测试前索引不存在
	if model.IndexExists() {
		if err := model.RemoveIndex(); err != nil {
			t.Fatalf("清理索引失败: %v", err)
		}
	}

	t.Run("Index 方法返回正确的索引名", func(t *testing.T) {
		expected := "article_index"
		actual := model.Index()
		if actual != expected {
			t.Errorf("期望索引名 %s, 实际得到 %s", expected, actual)
		}
	})

	t.Run("Mapping 方法正确读取映射文件", func(t *testing.T) {
		mapping := model.Mapping()
		if mapping == "" {
			t.Error("映射内容不能为空")
		}
		// 检查是否包含必要的 Elasticsearch 映射结构
		if !strings.Contains(mapping, `"mappings"`) || !strings.Contains(mapping, `"properties"`) {
			t.Error("映射文件缺少必要的结构")
		}
	})

	t.Run("索引不存在时 IndexExists 返回 false", func(t *testing.T) {
		if model.IndexExists() {
			t.Error("预期索引不存在，但 IndexExists 返回 true")
		}
	})

	t.Run("创建索引后能够正确存在", func(t *testing.T) {
		if err := model.CreateIndex(); err != nil {
			t.Fatalf("创建索引失败: %v", err)
		}
		defer func() {
			if err := model.RemoveIndex(); err != nil {
				t.Errorf("清理索引失败: %v", err)
			}
		}()

		if !model.IndexExists() {
			t.Error("创建索引后应存在，但 IndexExists 返回 false")
		}
	})

	t.Run("删除索引后应不存在", func(t *testing.T) {
		// 确保索引存在
		if err := model.CreateIndex(); err != nil {
			t.Fatalf("准备测试索引失败: %v", err)
		}

		if err := model.RemoveIndex(); err != nil {
			t.Fatalf("删除索引失败: %v", err)
		}

		if model.IndexExists() {
			t.Error("删除索引后应不存在，但 IndexExists 返回 true")
		}
	})

	t.Run("重复创建索引应覆盖已有索引", func(t *testing.T) {
		// 第一次创建
		if err := model.CreateIndex(); err != nil {
			t.Fatalf("第一次创建索引失败: %v", err)
		}
		defer func() {
			if err := model.RemoveIndex(); err != nil {
				t.Errorf("最终清理失败: %v", err)
			}
		}()

		// 第二次创建
		if err := model.CreateIndex(); err != nil {
			t.Fatalf("第二次创建索引失败: %v", err)
		}

		if !model.IndexExists() {
			t.Error("重复创建后索引应仍然存在")
		}
	})
}

func TestPathValidation(t *testing.T) {
	// 验证当前工作目录
	wd, _ := os.Getwd()
	t.Log("当前工作目录:", wd)

	// 验证配置文件路径
	if _, err := os.Stat("settings.yaml"); err != nil {
		t.Error("找不到配置文件")
	}

	// 验证映射文件路径
	mappingPath := "models/esmodels/article_mapper.json"
	if _, err := os.Stat(mappingPath); err != nil {
		t.Errorf("找不到映射文件: %s", mappingPath)
	}
}

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
