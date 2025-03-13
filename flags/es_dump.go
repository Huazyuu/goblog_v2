package flags

import (
	"backend/global"
	"backend/models/esmodels"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type ESRawMessage struct {
	Row json.RawMessage `json:"row"`
	ID  string          `json:"id"`
}

type ESIndexResponse struct {
	Data    []ESRawMessage `json:"data"`
	Mapping string         `json:"mapping"`
	Index   string         `json:"index"`
}

func esDump() {
	index := esmodels.FullTextModel{}.Index()
	mapping := esmodels.FullTextModel{}.Mapping()
	esDumpServer(index, mapping)
	index = esmodels.ArticleModel{}.Index()
	mapping = esmodels.ArticleModel{}.Mapping()
	esDumpServer(index, mapping)
}

func esDumpServer(index, mapping string) {
	res, err := global.ESClient.Search(index).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).Do(context.Background())

	if err != nil {
		logrus.Fatalf("%s err: %s", index, err.Error())
	}

	var dataList []ESRawMessage
	for _, hit := range res.Hits.Hits {
		dataList = append(dataList, ESRawMessage{
			Row: hit.Source,
			ID:  hit.Id,
		})
	}
	response := ESIndexResponse{
		Mapping: mapping,
		Index:   index,
		Data:    dataList,
	}

	fileName := fmt.Sprintf("./dump/es/%s_%s.json", index, time.Now().Format("20060102"))
	file, _ := os.Create(fileName)

	byteData, _ := json.Marshal(response)
	file.Write(byteData)
	file.Close()

	logrus.Infof("索引 %s 导出成功  %s", index, fileName)
}
