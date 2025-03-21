package menuService

import (
	"backend/controller/resp"
	"backend/repository/menu_banner_repo"
	"backend/repository/menu_repo"
)

func GetMenuDetail(id uint) (resp.MenuResponse, error) {
	menu, err := menu_repo.GetMenuByID(id)
	if err != nil {
		return resp.MenuResponse{}, err
	}
	menuBanners, err := menu_banner_repo.GetMenuBannersByID(menu.ID)
	if err != nil {
		return resp.MenuResponse{}, err
	}
	var banners = make([]resp.Banner, 0)
	for _, banner := range menuBanners {
		if menu.ID != banner.MenuID {
			continue
		}
		banners = append(banners, resp.Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := resp.MenuResponse{
		MenuModel: menu,
		Banners:   banners,
	}
	return menuResponse, nil
}
func GetMenuDetailByPath(path string) (resp.MenuResponse, error) {
	menu, err := menu_repo.GetMenuByPath(path)
	if err != nil {
		return resp.MenuResponse{}, err
	}
	menuBanners, err := menu_banner_repo.GetMenuBannersByID(menu.ID)
	if err != nil {
		return resp.MenuResponse{}, err
	}
	var banners = make([]resp.Banner, 0)
	for _, banner := range menuBanners {
		if menu.ID != banner.MenuID {
			continue
		}
		banners = append(banners, resp.Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := resp.MenuResponse{
		MenuModel: menu,
		Banners:   banners,
	}
	return menuResponse, nil
}
