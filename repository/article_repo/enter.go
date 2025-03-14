package article_repo

import (
	"backend/controller/req"
	"github.com/olivere/elastic/v7"
)

type Option struct {
	req.PageInfo
	Fields   []string
	Tag      string
	Category string
	Query    *elastic.BoolQuery
}

// GetForm Form 方法用于指定从第几个结果开始返回 常用于分页查询。
func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
