package cron

import (
	"backend/global"
	"github.com/robfig/cron/v3"
	"time"
)

func CronInit() {
	location, err := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(location))
	_, err = Cron.AddFunc("0 0 3 * * *", SyncArticleData)
	if err != nil {
		global.Log.Error(err)
	}
	_, err = Cron.AddFunc("0 0 3 * * *", SyncCommentData)
	if err != nil {
		global.Log.Error(err)
	}
	_, err = Cron.AddFunc("0 */5 * * * *", SyncLogs)
	if err != nil {
		global.Log.Error(err)
	}
	Cron.Start()
}
