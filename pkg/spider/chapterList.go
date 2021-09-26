package spider

import (
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"vtmtea.com/f.cli/model"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func ChapterList(link string, sourceModel model.SourceModel, cronId uint64) {
	//处理一下链接，校验是否包含domain
	if !strings.Contains(link, sourceModel.Domain) {
		link = sourceModel.Domain + link
	}
	bookExist := model.BookExist(link)
	if bookExist {
		return
	}
	c := colly.NewCollector()
	c.OnHTML("html", func(e *colly.HTMLElement) {
		authorName := e.ChildAttr(sourceModel.AuthorRoute, "content")
		categoryName := e.ChildAttr(sourceModel.CategoryRoute, "content")
		author, _ := model.GetAuthor(authorName)
		category, _ := model.GetCategoryByMap(categoryName)
		bookStatus := 1
		if strings.Contains(e.ChildAttr(sourceModel.StatusRoute, "content"), "连载") {
			bookStatus = 0
		}
		bookModel := model.BookModel{
			Title:           e.ChildAttr(sourceModel.BookNameRoute, "content"),
			AuthorID:        author.ID,
			CategoryID:      category.Id,
			CoverImage:      e.ChildAttr(sourceModel.CoverRoute, "content"),
			Description:     e.ChildAttr(sourceModel.DescriptionRoute, "content"),
			Rate:            0,
			RatePeopleCount: 0,
			ClickCount:      0,
			CollectCount:    0,
			RecommendCount:  0,
			LastUpdateTime:  e.ChildAttr(sourceModel.LastUdpateRoute, "content"),
			BookStatus:      bookStatus,
		}
		bookModel.Create()
		author.UpdateCount()
		category.UpdateCount()
		log.Infof("创建书本：%s", bookModel.Title)
		//log.Infof("插入章节列表")
		//i := 0
		//e.ForEach(sourceModel.ChapterListRoute, func(i int, element *colly.HTMLElement) {
		//	chapter := model.BookChapterModel{
		//		Title:      element.Text,
		//		SourceID:   sourceModel.Id,
		//		BookID:     bookModel.Id,
		//		SourceLink: element.Attr("href"),
		//		Content:    "",
		//	}
		//	chapter.Create()
		//	i+=1
		//})
		//log.Infof("章节列表插入完成, 共插入%d张", i)
		bookMapSource, err := model.GetSingleData(bookModel.Id, sourceModel.Id)
		if err != nil && err == gorm.ErrRecordNotFound {
			bookMapSource.SourceLink = link
			bookMapSource.Create()
		}
		if cronId > 0 {
			model.DeleteRecord(cronId)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	c.OnError(func(r *colly.Response, err error) {
		if cronId == 0 {
			m := model.FetchFailModel{
				SourceId:   sourceModel.Id,
				SourceLink: link,
			}
			m.Create()
		}
	})

	c.Visit(link)
}
