package flags

import (
	"backend/global"
	"backend/models/esmodels"
	"backend/service/esService/indexService"
)

func esCreate() {
	idxService := indexService.NewIndexService(global.ESClient, global.Log)
	err:=idxService.CreateIndexWithRetry(esmodels.ArticleModel{})
	if err != nil {
		global.Log.Error("CreateIndexWithRetry", "err", err)
		return
	}
	err=idxService.CreateIndexWithRetry(esmodels.FullTextModel{})
	if err != nil {
		global.Log.Error("CreateIndexWithRetry", "err", err)
		return
	}
}
