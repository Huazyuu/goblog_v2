package advertsService

import "C"
import (
	"backend/models/req"
	"backend/models/sqlmodels"
	"backend/repository/advert_repo"
	"errors"
)

func AdvertCreateService(cr req.AdvertRequest) (string, error) {
	var ad sqlmodels.AdvertModel
	ad, err := advert_repo.GetByTitle(cr.Title)
	if err == nil {
		return "广告已存在", errors.New("广告已存在")
	}
	ad.Title = cr.Title
	ad.Href = cr.Href
	ad.Images = cr.Images
	ad.IsShow = cr.IsShow
	if err = advert_repo.CreateAdvert(ad); err != nil {
		return "增加广告失败", errors.New("增加广告失败")
	}
	return "增加广告成功", nil
}
