package data_repo

import (
	"backend/models/esmodels"
	"backend/models/sqlmodels"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type DataCount struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	ChatGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
}

type DataService struct {
	DB       *gorm.DB
	ESClient *elastic.Client
}

func NewDataService(db *gorm.DB, es *elastic.Client) *DataService {
	return &DataService{DB: db, ESClient: es}
}

func (s *DataService) GetCounts() (DataCount, error) {
	var (
		cnts    DataCount
		wg      sync.WaitGroup
		errChan = make(chan error, 7)
		ctx     = context.Background()
	)

	wg.Add(6)
	go s.getArticleCount(ctx, &wg, errChan, &cnts)
	go s.queryCount(&wg, errChan, sqlmodels.UserModel{}, &cnts.UserCount, "用户统计失败")
	go s.queryCount(&wg, errChan, sqlmodels.MessageModel{}, &cnts.MessageCount, "消息统计失败")
	go s.queryCount(&wg, errChan, sqlmodels.ChatModel{IsGroup: true}, &cnts.ChatGroupCount, "群聊统计失败")
	go s.queryDailyCount(&wg, errChan, sqlmodels.LoginDataModel{}, &cnts.NowLoginCount, "当日登录统计失败")
	go s.queryDailyCount(&wg, errChan, sqlmodels.UserModel{}, &cnts.NowSignCount, "当日注册统计失败")

	// 错误处理
	go func() {
		wg.Wait()
		close(errChan)
	}()

	var finalErr error
	for e := range errChan {
		if finalErr == nil {
			finalErr = e
		} else {
			finalErr = fmt.Errorf("%v; %w", finalErr, e)
		}
	}

	return cnts, finalErr
}

func (s *DataService) getArticleCount(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error, cnts *DataCount) {
	defer wg.Done()

	result, err := s.ESClient.Search(esmodels.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(ctx)

	if err != nil {
		errChan <- errors.Wrap(err, "文章统计失败")
		return
	}
	cnts.ArticleCount = int(result.Hits.TotalHits.Value)
}

func (s *DataService) queryCount(wg *sync.WaitGroup, errChan chan<- error, model interface{}, dest *int, errMsg string) {
	defer wg.Done()

	if err := s.DB.Model(model).Select("count(id)").Scan(dest).Error; err != nil {
		errChan <- errors.Wrap(err, errMsg)
	}
}

func (s *DataService) queryDailyCount(wg *sync.WaitGroup, errChan chan<- error, model interface{}, dest *int, errMsg string) {
	defer wg.Done()

	if err := s.DB.Model(model).
		Where("created_at >= CURDATE()").
		Select("count(id)").
		Scan(dest).Error; err != nil {
		errChan <- errors.Wrap(err, errMsg)
	}
}
