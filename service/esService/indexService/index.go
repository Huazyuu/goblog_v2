package indexService

import (
	"backend/models/esmodels"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type Service struct {
	esClient *elastic.Client
	logger   *logrus.Logger
}

// NewIndexService 构造函数（依赖注入）
func NewIndexService(esClient *elastic.Client, logger *logrus.Logger) *Service {
	return &Service{
		esClient: esClient,
		logger:   logger,
	}
}

// CreateIndexWithRetry 带重试机制的索引创建
func (s *Service) CreateIndexWithRetry( ei esmodels.ESIndexInterFace) error {
	if s.esClient == nil {
		return fmt.Errorf("elasticsearch client not initialized")
	}

	if err := ei.CreateIndex(); err != nil {
		return err
	}

	return nil
}
