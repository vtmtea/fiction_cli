package book

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "vtmtea.com/f.cli/handler"
	"vtmtea.com/f.cli/model"
	"vtmtea.com/f.cli/pkg/errno"
	"vtmtea.com/f.cli/pkg/spider"
)

func UpdateChapterList(c *gin.Context) {
	bookId := c.DefaultQuery("bookId", "0")
	sourceId := c.DefaultQuery("sourceId", "0")
	bookIdToInt, err := strconv.Atoi(bookId)
	if err != nil {
		SendResponse(c, errno.ErrGetBookId, nil)
		return
	}
	sourceIdToInt, err := strconv.Atoi(sourceId)
	if err != nil {
		SendResponse(c, errno.ErrGetSourceId, nil)
		return
	}
	bookModel := model.GetBookById(uint64(bookIdToInt))
	sourceModel, _ := model.GetSourceById(uint64(sourceIdToInt))
	go spider.ChapterList(bookModel, sourceModel)
	SendResponse(c, nil, nil)
}
