package cron

import (
	"backend/global"
	"backend/models/sqlmodels"
	"backend/service/redisService"
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

// 创建批量处理通道
type updateTask struct {
	key   string
	count int
}

// SyncCommentData 优化后的评论数据同步方法
func SyncCommentData() {
	ctx := context.Background()
	// 获取 Redis 数据
	commentDiggInfo := redisService.NewCommentDigg().GetAll()
	if len(commentDiggInfo) == 0 {
		global.Log.Info("无评论点赞数据需要同步")
		return
	}

	taskCh := make(chan updateTask, 100)
	errCh := make(chan error, 1)

	// 启动处理协程
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := processBatchUpdates(ctx, taskCh); err != nil {
			select {
			case errCh <- err:
			default:
			}
		}
	}()

	// 分发任务
	sendTasks(ctx, commentDiggInfo, taskCh)

	// 等待处理完成
	close(taskCh)
	wg.Wait()

	// 错误处理
	select {
	case err := <-errCh:
		global.Log.Errorf("批量更新出现错误: %v", err)
		return
	default:
	}

	// 清除 Redis 数据（仅在成功时）
	redisService.NewCommentDigg().Clear()
	return
}

// 分发任务到通道
func sendTasks(ctx context.Context, data map[string]int, ch chan<- updateTask) {
	for key, count := range data {
		select {
		case ch <- updateTask{key: key, count: count}:
		case <-ctx.Done():
			return
		}
	}
}

// 批量处理更新
func processBatchUpdates(ctx context.Context, taskCh <-chan updateTask) error {
	const batchSize = 100
	var (
		batch      = make(map[string]int, batchSize)
		totalCount int
	)

	for {
		select {
		case task, ok := <-taskCh:
			if !ok {
				// 处理剩余批次
				if len(batch) > 0 {
					if err := executeBatchUpdate(ctx, batch); err != nil {
						return err
					}
				}
				global.Log.Infof("共处理 %d 条评论点赞更新", totalCount)
				return nil
			}

			batch[task.key] = task.count
			totalCount++

			// 达到批次大小时执行更新
			if len(batch) >= batchSize {
				if err := executeBatchUpdate(ctx, batch); err != nil {
					return err
				}
				batch = make(map[string]int, batchSize)
			}

		case <-ctx.Done():
			global.Log.Warn("同步操作被取消，剩余数据将不会提交")
			return ctx.Err()
		}
	}
}

// 执行批量更新
func executeBatchUpdate(ctx context.Context, batch map[string]int) error {
	// 构建 CASE WHEN 语句
	var (
		caseStmt strings.Builder
		args     []interface{}
		keys     []string
	)

	caseStmt.WriteString("CASE id ")
	for key, count := range batch {
		caseStmt.WriteString("WHEN ? THEN digg_count + ? ")
		args = append(args, key, count)
		keys = append(keys, key)
	}
	caseStmt.WriteString("END")

	// 在事务中执行更新
	err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 批量更新
		result := tx.Model(&sqlmodels.CommentModel{}).
			Where("id IN (?)", keys).
			Update("digg_count", gorm.Expr(caseStmt.String(), args...))

		if result.Error != nil {
			return result.Error
		}

		// 验证更新数量
		if int(result.RowsAffected) != len(keys) {
			global.Log.Warnf("更新数量不匹配，预期 %d 实际 %d", len(keys), result.RowsAffected)
		}

		return nil
	})

	// 错误重试逻辑
	if err != nil {
		if shouldRetry(err) {
			global.Log.Warnf("数据库更新失败，尝试重试: %v", err)
			time.Sleep(1 * time.Second)
			return executeBatchUpdate(ctx, batch)
		}
		return fmt.Errorf("批量更新失败: %w", err)
	}

	return nil
}

// 可重试错误判断
func shouldRetry(err error) bool {
	// 可根据具体数据库错误类型扩展
	if strings.Contains(err.Error(), "deadlock") ||
		strings.Contains(err.Error(), "timeout") ||
		strings.Contains(err.Error(), "busy") {
		return true
	}
	return false
}
