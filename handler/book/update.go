package book

import (
	"github.com/gin-gonic/gin"
	. "vtmtea.com/f.cli/handler"
	"vtmtea.com/f.cli/pkg/errno"
)

type Request struct {
	BookId   string `json:"book_id"`
	SourceId string `json:"source_id"`
}

func Update(c *gin.Context) {
	var err error
	var r Request
	if err = c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
}
