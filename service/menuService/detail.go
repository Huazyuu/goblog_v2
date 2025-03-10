package menuService

import (
	"backend/repository/menu_banner_repo"
	"backend/repository/menu_repo"
)

func GetMenuDetail(id uint) (MenuResponse, error) {
	menu, err := menu_repo.GetMenuByID(id)
	if err != nil {
		return MenuResponse{}, err
	}
	menuBanners, err := menu_banner_repo.GetMenuBannersByID(menu.ID)
	if err != nil {
		return MenuResponse{}, err
	}
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menu.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menu,
		Banners:   banners,
	}
	return menuResponse, nil
}
func GetMenuDetailByPath(path string) (MenuResponse, error) {
	menu, err := menu_repo.GetMenuByPath(path)
	if err != nil {
		return MenuResponse{}, err
	}
	menuBanners, err := menu_banner_repo.GetMenuBannersByID(menu.ID)
	if err != nil {
		return MenuResponse{}, err
	}
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menu.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menu,
		Banners:   banners,
	}
	return menuResponse, nil
}
