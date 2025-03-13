package flags

import (
	"backend/global"
	"backend/models/esmodels"
	"backend/service/esService/indexService"
)

func esCreate() {
	idxService := indexService.NewIndexService(global.ESClient, global.Log)
	err := idxService.CreateIndex(esmodels.ArticleModel{})
	if err != nil {
		global.Log.Error("CreateIndex", "err", err)
		return
	}
	err = idxService.CreateIndex(esmodels.FullTextModel{})
	if err != nil {
		global.Log.Error("CreateIndex", "err", err)
		return
	}
}
