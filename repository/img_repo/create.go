package img_repo

import (
	"backend/global"
	"backend/models/sqlmodels"
)

func CreateBanner(model sqlmodels.BannerModel) (err error) {
	return global.DB.Create(&model).Error
}
