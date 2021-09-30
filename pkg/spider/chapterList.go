package spider

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"vtmtea.com/f.cli/model"
)

func ChapterList(bookModel model.BookModel, sourceModel model.SourceModel) {
	bookMapSource, err := model.GetSingleData(bookModel.Id, sourceModel.Id)
	if err != nil {
		log.WithFields(log.Fields{"bookId": bookModel.Id, "sourceId": sourceModel.Id}).Errorf("抓取小说章节列表失败；错误信息: %s", err.Error())
		return
	}

	chapterCount := model.GetChapterCount(sourceModel.Id, bookModel.Id)

	c := colly.NewCollector()

	c.OnHTML(sourceModel.ChapterListRoute, func(e *colly.HTMLElement) {
		if e.Index > chapterCount-1 {
			chapter := model.BookChapterModel{
				Title:      e.Text,
				SourceID:   sourceModel.Id,
				BookID:     bookModel.Id,
				SourceLink: e.Attr("href"),
				Content:    "",
			}
			chapter.Create()
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.WithFields(log.Fields{"bookId": bookModel.Id, "sourceId": sourceModel.Id}).Errorf("抓取小说章节列表失败；错误信息: %s", err.Error())
	})

	c.Visit(bookMapSource.SourceLink)
}
