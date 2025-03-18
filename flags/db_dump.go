package flags

import (
	"backend/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

func dbDump() {
	mysql := global.Config.Mysql
	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("./dump/sql/%s_%s.sql", mysql.DB, timer)

	// 调用系统命令，执行 mysqldump 进行数据库导出，添加 -P 选项指定端口号
	fmt.Println(fmt.Sprintf("mysqldump -u%s -p%s -P%d -h%s --column-statistics=0 %s > %s", mysql.User, mysql.Password, mysql.Port, mysql.Host, mysql.DB, sqlPath))
	cmder := fmt.Sprintf("mysqldump -u%s -p%s -P%d -h%s %s > %s", mysql.User, mysql.Password, mysql.Port, mysql.Host, mysql.DB, sqlPath)
	cmd := exec.Command("sh", "-c", cmder)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logrus.Errorln(err.Error(), stderr.String())
		return
	}
	logrus.Infof("sql文件 %s 导出成功", sqlPath)
}
