package cron

import (
	"backend/global"
	"backend/models/esmodels"
	"backend/service/redisService"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"io"
	"sync"
	"time"
)

func SyncArticleData() {
	ctx := context.Background()
	// 1. 创建批量处理器
	bp, err := global.ESClient.BulkProcessor().
		Name("ShutdownSync").
		Workers(4).
		BulkActions(500). // 每500个操作提交一次
		FlushInterval(1 * time.Second).
		Do(ctx)
	if err != nil {
		global.Log.Error(err)
		return
	}
	defer bp.Close()

	articleDiggMap := redisService.NewArticleDigg().GetAll()
	articleLookMap := redisService.NewArticleLook().GetAll()
	commentMap := redisService.NewCommentCount().GetAll()

	// 并行处理ES数据
	var wg sync.WaitGroup
	workerCount := 5
	docCh := make(chan *elastic.SearchHit, 1000)

	// 启动工作池
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for hit := range docCh {
				processDocument(hit, bp, articleDiggMap, articleLookMap, commentMap)
			}
		}()
	}

	// 滚动查询ES数据
	scroll := global.ESClient.Scroll(esmodels.ArticleModel{}.Index()).
		Size(1000).
		Query(elastic.NewMatchAllQuery()).
		KeepAlive("2m")

	for {
		res, err := scroll.Do(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			global.Log.Error(err)

			return
		}

		// 分发文档到工作池
		for _, hit := range res.Hits.Hits {
			select {
			case docCh <- hit:
			case <-ctx.Done():
				return
			}
		}
	}

	close(docCh)
	wg.Wait()

	// 最后刷新剩余操作
	if err = bp.Flush(); err != nil {
		global.Log.Error(err)

		return
	}

	// 6. 清理Redis数据（保持原始接口）
	redisService.NewArticleDigg().Clear()
	redisService.NewArticleLook().Clear()
	redisService.NewCommentCount().Clear()

	return
}

func processDocument(hit *elastic.SearchHit, bp *elastic.BulkProcessor, diggMap, lookMap, commentMap map[string]int) {

	var article esmodels.ArticleModel
	if err := json.Unmarshal(hit.Source, &article); err != nil {
		global.Log.Errorf("文档解析失败 ID:%s - %v", hit.Id, err)
		return
	}

	// 计算增量
	digg := diggMap[hit.Id]
	look := lookMap[hit.Id]
	comment := commentMap[hit.Id]

	if digg == 0 && look == 0 && comment == 0 {
		return
	}

	// 创建批量更新请求
	updateReq := elastic.NewBulkUpdateRequest().
		Index(esmodels.ArticleModel{}.Index()).
		Id(hit.Id).
		Doc(map[string]interface{}{
			"digg_count":    article.DiggCount + digg,
			"look_count":    article.LookCount + look,
			"comment_count": article.CommentCount + comment,
		}).
		RetryOnConflict(3)

	bp.Add(updateReq)
}
