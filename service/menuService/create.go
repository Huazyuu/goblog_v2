package menuService

import (
	"backend/controller/req"
	"backend/global"
	"backend/models/sqlmodels"
	"backend/repository/img_repo"
	"backend/repository/menu_banner_repo"
	"backend/repository/menu_repo"
	"errors"
	"fmt"
)

func MenuCreateService(cr req.MenuRequest) (string, error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查重复
	exist, err := menu_repo.CheckMenuDuplicate(tx, cr.Title, cr.Path)
	if err != nil {
		tx.Rollback()
		return "系统错误", fmt.Errorf("重复检查失败: %w", err)
	}
	if exist {
		tx.Rollback()
		return "菜单已存在", errors.New("duplicate menu")
	}

	// 创建菜单
	menu := &sqlmodels.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	if err = menu_repo.CreateMenu(tx, menu); err != nil {
		tx.Rollback()
		global.Log.Error("创建菜单失败: " + err.Error())
		return "系统错误", fmt.Errorf("创建菜单失败: %w", err)
	}

	// 处理图片关联
	if len(cr.ImageSortList) > 0 {
		// 收集所有图片ID
		var bannerIDs []uint
		for _, img := range cr.ImageSortList {
			bannerIDs = append(bannerIDs, img.ImageID)
		}

		// 验证图片有效性
		invalidIDs, err := img_repo.GetInvalidBannerIDs(tx, bannerIDs)
		if err != nil {
			tx.Rollback()
			return "系统错误", fmt.Errorf("图片验证失败: %w", err)
		}

		if len(invalidIDs) > 0 {
			tx.Rollback()
			return fmt.Sprintf("无效图片ID: %v", invalidIDs),
				errors.New("invalid banner ids")
		}

		// 创建关联关系
		var menuBanners []sqlmodels.MenuBannerModel
		for _, img := range cr.ImageSortList {
			menuBanners = append(menuBanners, sqlmodels.MenuBannerModel{
				MenuID:   menu.ID,
				BannerID: img.ImageID,
				Sort:     img.Sort,
			})
		}

		if err = menu_banner_repo.CreateMenuBanners(tx, menuBanners); err != nil {
			tx.Rollback()
			global.Log.Error("创建关联失败: " + err.Error())
			return "系统错误", fmt.Errorf("创建关联失败: %w", err)
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		global.Log.Error("事务提交失败: " + err.Error())
		return "系统错误", fmt.Errorf("事务提交失败: %w", err)
	}

	return "菜单创建成功", nil
}
