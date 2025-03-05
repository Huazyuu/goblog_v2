package settingsService

import (
	"backend/core"
	"backend/global"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
)

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(core.ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
