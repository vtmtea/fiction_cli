package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"vtmtea.com/f.cli/model"
)

func List(link string, sourceRuleModel model.SourceRuleModel, sourceModel model.SourceModel) {
	log.Infof("准备访问链接: %s", link)
	c := colly.NewCollector()
	c.OnHTML(sourceRuleModel.BookListRoute, func(e *colly.HTMLElement) {
		bookName := e.ChildText(sourceRuleModel.BookNameRoute)
		authorName := e.ChildText(sourceRuleModel.AuthorRoute)
		bookLink := e.ChildAttr(sourceRuleModel.BookLinkRoute, "href")

		author, err := model.GetAuthor(authorName)
		if err != nil && err == gorm.ErrRecordNotFound {
			log.Infof("准备创建作者： %s", author.Name)
			author.Create()
			go ChapterInfo(bookLink, sourceModel, 0)
			return
		}
		_, err = model.GetBookByMapAttr(map[string]interface{}{"title": bookName, "status": 1, "author_id": author.ID})
		if err != nil && err == gorm.ErrRecordNotFound {
			go ChapterInfo(bookLink, sourceModel, 0)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorf("visit %s get error", r.Request.URL)
	})

	c.Visit(link)
}
