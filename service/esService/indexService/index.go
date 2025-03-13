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

// CreateIndex 索引创建
func (s *Service) CreateIndex(ei esmodels.ESIndexInterFace) error {
	if s.esClient == nil {
		s.logger.Warn("CreateIndex", "ES client not initialized")
		return fmt.Errorf("elasticsearch client not initialized")
	}

	if err := ei.CreateIndex(); err != nil {
		s.logger.Warn("CreateIndex", "CreateIndex err", err)
		return err
	}

	return nil
}
