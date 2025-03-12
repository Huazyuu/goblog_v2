package article_repo

import (
	"backend/global"
	"backend/models/esmodels"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func CreateArticle(article *esmodels.ArticleModel) error {
	indexResponse, err := global.ESClient.Index().
		Index(article.Index()).
		BodyJson(article).Refresh("true").Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	article.ID = indexResponse.Id
	return nil
}

func ISExistData(article esmodels.ArticleModel) bool {
	res, err := global.ESClient.
		Search(article.Index()).
		Query(elastic.NewTermQuery("keyword", article.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}
func GetDataByID(id string) (article esmodels.ArticleModel, err error) {
	res, err := global.ESClient.
		Get().
		Index(article.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return article, err
	}
	err = json.Unmarshal(res.Source, &article)
	return article, err
}
