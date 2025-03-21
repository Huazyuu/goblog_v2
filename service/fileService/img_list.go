package fileService

import (
	"backend/controller/req"
	"backend/controller/resp"
	"backend/models/sqlmodels"
	"backend/repository/article_repo"
)

func ImageListService(cr req.PageInfo) ([]resp.ImageListResponse, error) {
	imageList, _, err := req.ComList(sqlmodels.BannerModel{}, req.Option{
		PageInfo: cr,
		Debug:    true,
		Likes:    []string{"name", "path"},
		Preload:  []string{"MenusBanner"},
	})
	if err != nil {
		return nil, err
	}

	var imageIDList []interface{}
	for _, model := range imageList {
		imageIDList = append(imageIDList, model.ID)
	}

	var imageIDArticleCountMap = map[uint]int{}
	bannerids, err := article_repo.GetArticleBannerID(imageIDList)
	if err != nil {
		return nil, err
	}

	for _, bannerid := range bannerids {
		val, ok := imageIDArticleCountMap[bannerid]
		if !ok {
			imageIDArticleCountMap[bannerid] = 1
		} else {
			imageIDArticleCountMap[bannerid] = val + 1
		}
	}

	var list = make([]resp.ImageListResponse, 0)
	for _, model := range imageList {
		list = append(list, resp.ImageListResponse{
			BannerModel:  model,
			BannerCount:  len(model.MenusBanner),
			ArticleCount: imageIDArticleCountMap[model.ID],
		})
	}
	return list, nil
}
