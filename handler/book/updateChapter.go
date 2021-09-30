package book

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "vtmtea.com/f.cli/handler"
	"vtmtea.com/f.cli/model"
	"vtmtea.com/f.cli/pkg/errno"
	"vtmtea.com/f.cli/pkg/spider"
)

func UpdateChapter(c *gin.Context) {
	bookChapterId := c.DefaultQuery("bookChapterId", "0")
	sourceId := c.DefaultQuery("sourceId", "0")
	bookChapterIdToInt, err := strconv.Atoi(bookChapterId)
	if err != nil {
		SendResponse(c, errno.ErrGetBookChapterId, nil)
		return
	}
	sourceIdToInt, err := strconv.Atoi(sourceId)
	if err != nil {
		SendResponse(c, errno.ErrGetSourceId, nil)
		return
	}
	bookChapterModel := model.GetChapterById(uint64(bookChapterIdToInt))
	sourceModel, _ := model.GetSourceById(uint64(sourceIdToInt))
	bookMapSourceModel, _ := model.GetSingleData(bookChapterModel.BookID, uint64(sourceIdToInt))
	go spider.Chapter(bookMapSourceModel.SourceLink+bookChapterModel.SourceLink, bookChapterModel, sourceModel)
	SendResponse(c, nil, nil)
}
