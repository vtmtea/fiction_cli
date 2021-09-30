package book

import (
	"github.com/gin-gonic/gin"
	. "vtmtea.com/f.cli/handler"
	"vtmtea.com/f.cli/model"
	"vtmtea.com/f.cli/pkg/errno"
	"vtmtea.com/f.cli/pkg/spider"
)

type postBody struct {
	SourceLink string `json:"sourceLink"`
	SourceId   uint64 `json:"sourceId" binding:"required"`
}

func FetchOne(c *gin.Context) {
	var post postBody
	if err := c.ShouldBindJSON(&post); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	sourceModel, _ := model.GetSourceById(post.SourceId)
	go spider.ChapterInfo(post.SourceLink, sourceModel, 0)
	SendResponse(c, nil, nil)
}
