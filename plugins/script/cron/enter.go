package cron

import (
	"backend/global"
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithLocation(timezone))
	_, err := Cron.AddFunc("0 0 * * *", SyncArticleData)
	if err != nil {
		global.Log.Error(err)
	}
	_, err = Cron.AddFunc("0 0 * * *", SyncCommentData)
	if err != nil {
		global.Log.Error(err)
	}
	_, err = Cron.AddFunc("*/10 * * * *", SyncLogs)
	if err != nil {
		global.Log.Error(err)
	}
	Cron.Start()
}
