package redisService

import (
	"backend/global"
	"backend/models/sqlmodels"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const (
	logQueueKey         = "logstash:queue" // 日志队列KEY
	maxBatchSize        = 100              // 最大批量处理条数
	expireDuration      = 24 * time.Hour   // Redis数据过期时间
	bufferFlushInterval = 5 * time.Minute  // 批量处理间隔
)

type RedisStore struct {
	client *redis.Client
	db     *gorm.DB
	buffer []*sqlmodels.LogModel
}

func NewRedisStore(client *redis.Client, db *gorm.DB) *RedisStore {
	return &RedisStore{
		client: client,
		db:     db,
		buffer: make([]*sqlmodels.LogModel, 0, maxBatchSize),
	}
}

// AddLog 添加日志到缓冲并触发批量保存
func (rs *RedisStore) AddLog(log *sqlmodels.LogModel) error {
	rs.buffer = append(rs.buffer, log)

	if len(rs.buffer) >= maxBatchSize {
		return rs.Flush()
	}
	return nil
}

// Flush 批量保存到Redis（使用RPush确保顺序）
func (rs *RedisStore) Flush() error {
	pipe := rs.client.Pipeline()

	for _, log := range rs.buffer {
		data, _ := json.Marshal(log)
		pipe.RPush(logQueueKey, data) // 修改为RPush
	}

	// 设置过期时间
	pipe.Expire(logQueueKey, expireDuration)

	_, err := pipe.Exec()
	if err == nil {
		rs.buffer = rs.buffer[:0] // 清空缓冲
	}
	return err
}

// GetLogs 从Redis获取批量日志（从队列头部获取最早的数据）
func (rs *RedisStore) GetLogs(batchSize int64) ([]*sqlmodels.LogModel, error) {
	r := rs.client.LRange(logQueueKey, 0, batchSize-1)
	result, err := r.Result()
	global.Log.Info("result:", result)
	if err != nil {
		return nil, err
	}

	var logs []*sqlmodels.LogModel
	for _, item := range result {
		var log sqlmodels.LogModel
		if err = json.Unmarshal([]byte(item), &log); err != nil {
			global.Log.Errorf(err.Error())
		}
		logs = append(logs, &log)
	}

	return logs, nil
}

// RemoveLogs 删除已处理的日志（保留batchSize之后的元素）
func (rs *RedisStore) RemoveLogs(batchSize int64) error {
	return rs.client.LTrim(logQueueKey, batchSize, -1).Err()
}
