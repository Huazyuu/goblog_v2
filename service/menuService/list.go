package menuService

import (
	"backend/controller/resp"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/repository/menu_banner_repo"
	"backend/repository/menu_repo"
)

func GetFullMenuList() ([]resp.MenuResponse, error) {
	// 获取菜单列表
	menus, err := menu_repo.GetMenuList()
	if err != nil {
		return nil, err
	}

	// 获取关联的菜单横幅数据
	menuIDs := make([]uint, 0, len(menus))
	for _, menu := range menus {
		menuIDs = append(menuIDs, menu.ID)
	}

	menuBanners, err := menu_banner_repo.GetMenuBannersByIDList(menuIDs)
	if err != nil {
		return nil, err
	}

	menuBannerMap := make(map[uint][]resp.Banner)
	for _, mb := range menuBanners {
		menuBannerMap[mb.MenuID] = append(menuBannerMap[mb.MenuID], resp.Banner{
			ID:   mb.BannerID,
			Path: mb.BannerModel.Path,
		})
	}

	// 组装响应数据
	response := make([]resp.MenuResponse, 0, len(menus))
	for _, menu := range menus {
		response = append(response, resp.MenuResponse{
			MenuModel: menu,
			Banners:   menuBannerMap[menu.ID],
		})
	}

	return response, nil
}

func GetMenuNameList() (menuNameList []resp.MenuNameResponse, err error) {
	result := global.DB.Model(sqlmodels.MenuModel{}).Select("id", "title", "path")
	result.Scan(&menuNameList)
	err = result.Error
	if err != nil {
		return nil, err
	}
	return menuNameList, nil
}
