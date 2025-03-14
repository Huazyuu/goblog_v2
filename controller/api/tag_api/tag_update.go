package tag_api

import (
	"backend/controller/req"
	"backend/controller/res"
	"backend/service/tagService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr req.TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	uint64id, _ := strconv.ParseUint(id, 10, 64)
	uintid := uint(uint64id)
	msg, err := tagService.TagUpdate(uintid, cr)
	if err != nil {
		res.FailWithMessage(msg, c)
		return
	}
	res.OkWithMessage(msg, c)
}
