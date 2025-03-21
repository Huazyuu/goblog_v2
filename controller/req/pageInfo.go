package req

import (
	"backend/global"
	"fmt"
	"gorm.io/gorm"
)

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type Option struct {
	PageInfo          // 分页查询
	Likes    []string // 需要模糊匹配的字段列表
	Debug    bool     // 是否打印sql
	Where    *gorm.DB // 额外的查询
	Preload  []string // 预加载的字段列表
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	query := global.DB.Where(model)
	if option.Debug {
		query = query.Debug()
	}

	// 默认按照时间往后排
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	// 默认一页显示10条
	if option.Limit == 0 {
		option.Limit = 10
	}
	// 如果有高级查询就加上
	if option.Where != nil {
		query.Where(option.Where)
	}

	// 模糊匹配
	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			// 第一个模糊匹配和前面的关系是and关系，后面的和前面的模糊匹配是or的关系
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		// 整个模糊匹配它是一个整体
		query = query.Where(likeQuery)
	}

	// 查列表，获取总数
	count = query.Find(&list).RowsAffected

	// 预加载
	/*	type BannerModel struct {
			MenusBanner []MenuBannerModel `gorm:"foreignKey:BannerID" json:"-"`
		}
		type MenuBannerModel struct {
			BannerID int
		}
	*/
	// Preload 方法里的参数名称要和结构体字段名的大小写保持一致，
	// 也就是 "MenusBanner" 必须和 BannerModel 结构体中的 MenusBanner 字段名大小写相同
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	// 计算偏移
	offset := (option.Page - 1) * option.Limit

	err = query.Limit(option.Limit).
		Offset(offset).
		Order(option.Sort).Find(&list).Error

	return
}
