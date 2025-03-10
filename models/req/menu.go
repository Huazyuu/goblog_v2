package req

import "backend/models/diverseType"

type MenuRequest struct {
	Title         string            `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string            `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string            `json:"slogan" structs:"slogan"`
	Abstract      diverseType.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int               `json:"abstract_time" structs:"abstract_time"` // 切换的时间，单位秒
	BannerTime    int               `json:"banner_time" structs:"banner_time"`     // 切换的时间，单位秒
	Sort          int               `json:"sort" structs:"sort"`                   // 菜单的序号
	ImageSortList []ImageSort       `json:"image_sort_list" structs:"-"`           // 具体图片的顺序
}
type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
