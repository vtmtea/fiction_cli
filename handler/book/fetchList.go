package book

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	. "vtmtea.com/f.cli/handler"
	"vtmtea.com/f.cli/model"
	"vtmtea.com/f.cli/pkg/spider"
)

func FetchList(c *gin.Context) {
	sourceId := c.Query("sourceId")
	pageCount := c.DefaultQuery("pageCount", "1")
	sourceIdInt, _ := strconv.Atoi(sourceId)
	page, _ := strconv.Atoi(pageCount)
	sourceModel, _ := model.GetSourceById(uint64(sourceIdInt))
	sourceRuleModels, _ := model.GetSourceRule(uint64(sourceIdInt))
	for _, sourceRuleModel := range sourceRuleModels {
		if strings.Contains(sourceRuleModel.ListLink, "[page]") {
			for i := 1; i <= page; i++ {
				go spider.List(sourceModel.Domain+strings.ReplaceAll(sourceRuleModel.ListLink, "[page]", strconv.Itoa(i)), sourceRuleModel, sourceModel)
			}
		}
	}
	SendResponse(c, nil, nil)
}
