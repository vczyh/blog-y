package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreatePostAPI(c *gin.Context) (interface{}, error) {
	if err := AuthBaseAPI(c); err != nil {
		return nil, err
	}

	var req PostReq
	if err := c.Bind(&req); err != nil {
		return nil, fmt.Errorf("param error")
	}
	return CreatePostService(req)
}

func PostListAPI(c *gin.Context) (interface{}, error) {
	format := c.Query("format")
	tagId := c.Query("tag_id")

	switch format {
	case "list":
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		return PostListService(page, pageSize, tagId)
	case "timeline":
		return TimelineService()
	default:
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		return PostListService(page, pageSize, tagId)
	}

}

func GetPostByPostIdAPI(c *gin.Context) (interface{}, error) {
	postId := c.Param("postId")
	return GetPostByPostIdService(postId)
}

func UpdatePostAPI(c *gin.Context) (interface{}, error) {
	if err := AuthBaseAPI(c); err != nil {
		return nil, err
	}

	postId := c.Param("postId")
	var req PostReq
	if err := c.Bind(&req); err != nil {
		return nil, fmt.Errorf("param error")
	}
	return nil, UpdatePostService(postId, req)
}

func AddPostTagAPI(c *gin.Context) (interface{}, error) {
	if err := AuthBaseAPI(c); err != nil {
		return nil, err
	}

	postId := c.Param("postId")
	var req = struct {
		TagIds []string `json:"tag_ids"`
	}{}
	if err := c.Bind(&req); err != nil {
		l.Error("param error")
	}
	return nil, AddPostTagService(postId, req.TagIds)
}

func DeletePostTagAPI(c *gin.Context) (interface{}, error) {
	if err := AuthBaseAPI(c); err != nil {
		return nil, err
	}

	postId := c.Param("postId")
	tagId := c.Query("tag_id")
	return nil, DeletePostTagService(postId, []string{tagId})
}

func DeletePostAPI(c *gin.Context) (interface{}, error) {
	if err := AuthBaseAPI(c); err != nil {
		return nil, err
	}

	postId := c.Param("postId")
	return nil, DeletePostService(postId)
}
