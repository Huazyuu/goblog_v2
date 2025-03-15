package redisService

const (
	articleLookCountPrefix    = "article_look_count"    // 文章浏览量
	articleDiggCountPrefix    = "article_digg_count"    // 文章点赞数
	articleCommentCountPrefix = "article_comment_count" // 文章评论数
	commentDiggCountPrefix    = "comment_digg_count"    // 评论点赞数
)

func NewArticleDigg() RedisCount {
	return CountDB{
		Index: articleDiggCountPrefix,
	}
}
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookCountPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentCountPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggCountPrefix,
	}
}
