package spider

import (
	"vtmtea.com/f.cli/model"
)

func Cron() {
	fetFails := model.GetFailList(200)
	for _, failModel := range fetFails {
		sourceModel, _ := model.GetSourceById(failModel.SourceId)
		go ChapterList(failModel.SourceLink, sourceModel, failModel.Id)
	}
}
