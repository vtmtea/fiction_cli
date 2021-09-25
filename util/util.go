package util

import (
	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids/v2"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func HashIdInstance() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = viper.GetString("hash_salt")
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	return h
}

func HashIdEncode(id int) string {
	h := HashIdInstance()
	e, err := h.Encode([]int{id})
	if err != nil {
		return ""
	}
	return e
}

func HashIdDecode(hashId string) int {
	h := HashIdInstance()
	id, err := h.DecodeWithError(hashId)
	if err != nil {
		return 0
	}
	return id[0]
}
