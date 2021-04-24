package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateTagAPI(c *gin.Context) (interface{}, error) {
	if err := AuthBaseAPI(c); err != nil {
		return nil, err
	}

	var req TagReq
	if err := c.Bind(&req); err != nil {
		return nil, fmt.Errorf("param error")
	}
	return CreateTagService(req)
}

func TagListAPI(c *gin.Context) (interface{}, error) {
	return TagListService()
}
