package cron

import (
	"backend/global"
	"backend/service/redisService"
)

func SyncLogs() {
	log := redisService.NewRedisStore(global.Redis, global.DB)
	global.Log.Info("sync logs.....")
	processLogs(log)
}

func processLogs(rs *redisService.RedisStore) error {
	logs, err := rs.GetLogs(int64(100))
	if err != nil {
		return err
	}
	if len(logs) == 0 {
		return nil
	}

	// 使用事务批量插入
	tx := global.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(logs).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 删除已处理的数据
	if err := rs.RemoveLogs(int64(len(logs))); err != nil {
		return err
	}

	return nil
}
