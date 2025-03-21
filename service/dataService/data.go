package dataService

import (
	"backend/global"
	"backend/repository/data_repo"
)

func DataSumService() (counts data_repo.DataCount, err error) {
	counts, err = data_repo.NewDataService(global.DB, global.ESClient).GetCounts()
	if err != nil {
		return
	}
	return
}
