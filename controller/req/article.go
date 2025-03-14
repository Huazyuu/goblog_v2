package req

type ArticleRequest struct {
	Title    string   `json:"title" binding:"required" msg:"文章标题必填"`   // 文章标题
	Abstract string   `json:"abstract"`                                // 文章简介
	Content  string   `json:"content" binding:"required" msg:"文章内容必填"` // 文章内容
	Category string   `json:"category"`                                // 文章分类
	Source   string   `json:"source"`                                  // 文章来源
	Link     string   `json:"link"`                                    // 原文链接
	BannerID uint     `json:"banner_id"`                               // 文章封面id
	Tags     []string `json:"tags"`                                    // 文章标签
}
type ArticleSearchRequest struct {
	PageInfo
	Tag      string `json:"tags" form:"tags"`
	Category string `json:"category" form:"category"`
	IsUser   bool   `json:"is_user" form:"is_user"` // 根据这个参数判断是否显示我收藏的文章列表
	Date     string `json:"date" form:"date"`       // 发布时间搜索
}
type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}
type ArticleUpdateRequest struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`     // 文章标题
	Abstract string   `json:"abstract"`  // 文章简介
	Content  string   `json:"content"`   // 文章内容
	Category string   `json:"category"`  // 文章分类
	Source   string   `json:"source"`    // 文章来源
	Link     string   `json:"link"`      // 原文链接
	BannerID uint     `json:"banner_id"` // 文章封面id
	Tags     []string `json:"tags"`      // 文章标签
}
