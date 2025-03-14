package tag_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/tagService"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagCreateView(c *gin.Context) {
	var cr req.TagRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	msg, err := tagService.TagCreate(cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
