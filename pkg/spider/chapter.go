package spider

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"vtmtea.com/f.cli/model"
)

func Chapter(link string, bookChapterModel model.BookChapterModel, sourceModel model.SourceModel) {
	c := colly.NewCollector()

	c.OnHTML(sourceModel.ChapterContentRoute, func(e *colly.HTMLElement) {
		bookChapterModel.Content = e.Text
		bookChapterModel.SaveContent()
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorln("访问内容链接失败", err.Error())
	})

	c.Visit(link)
}
