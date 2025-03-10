package menuService

import (
	"backend/global"
	"backend/models/req"
	"backend/models/sqlmodels"
	"errors"
	"fmt"
)

func MenuCreate(cr req.MenuRequest) (string, error) {
	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查菜单是否重复
	cnt, err := sqlmodels.MenuModel{}.IsDuplicate(tx, cr.Title, cr.Path)
	if err != nil {
		return "查询错误", err
	}
	if cnt > 0 {
		return "重复的菜单", errors.New("重复的菜单")
	}

	// 创建菜单
	menuModel := sqlmodels.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	if err = tx.Create(&menuModel).Error; err != nil {
		tx.Rollback()
		global.Log.Error("创建菜单失败: " + err.Error())
		return "创建菜单失败", err
	}

	// 没有图片直接提交事务
	if len(cr.ImageSortList) == 0 {
		tx.Commit()
		return "菜单添加成功", nil
	}

	// 检查图片是否存在
	var invalidImgIDs []uint
	for _, img := range cr.ImageSortList {
		var banner sqlmodels.BannerModel
		if err = tx.Where("id = ?", img.ImageID).First(&banner).Error; err != nil {
			invalidImgIDs = append(invalidImgIDs, img.ImageID)
		}
	}

	// 存在无效图片ID时回滚
	if len(invalidImgIDs) > 0 {
		tx.Rollback()
		errorMsg := fmt.Sprintf("关联图片不存在: %v", invalidImgIDs)
		global.Log.Error(errorMsg)
		return errorMsg, errors.New(errorMsg)
	}

	// 创建关联关系
	var menuBannerList []sqlmodels.MenuBannerModel
	for _, img := range cr.ImageSortList {
		menuBannerList = append(menuBannerList, sqlmodels.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: img.ImageID,
			Sort:     img.Sort,
		})
	}

	if err = (&sqlmodels.MenuBannerModel{}).CreateMenuBannerTX(tx, menuBannerList); err != nil {
		tx.Rollback()
		global.Log.Error("创建关联失败: " + err.Error())
		return "图片关联失败", err
	}

	// 提交事务
	tx.Commit()
	return "菜单添加成功", nil
}
