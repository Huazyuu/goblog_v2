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

func UpdateMenu(menuID uint, cr req.MenuRequest) (string, error) {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if _, err := menu_repo.GetMenuByID(menuID); err != nil {
		tx.Rollback()
		return "用户不存在", err
	}
	// 清理旧关联
	if err := menu_banner_repo.ClearByMenuID(tx, menuID); err != nil {
		tx.Rollback()
		return "清理旧关联失败", err
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
				MenuID:   menuID,
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
	// 更新基础信息
	updateMap := map[string]interface{}{
		"title":         cr.Title,
		"path":          cr.Path,
		"slogan":        cr.Slogan,
		"abstract":      cr.Abstract,
		"abstract_time": cr.AbstractTime,
		"banner_time":   cr.BannerTime,
		"sort":          cr.Sort,
	}

	if err := menu_repo.UpdateMenu(tx, menuID, updateMap); err != nil {
		tx.Rollback()
		return "更新菜单失败", err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		global.Log.Error("事务提交失败: " + err.Error())
		return "系统错误", fmt.Errorf("事务提交失败: %w", err)
	}

	return "菜单更新成功", nil
}
