package route

import "github.com/gin-gonic/gin"

const (
	_success = 0
	_failed  = 1
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HandlerFunc func(*gin.Context) (interface{}, error)

func Handle(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handlerFunc(c)

		var res Res
		if err != nil {
			res.Code = _failed
			res.Message = err.Error()
		} else {
			res.Code = _success
			res.Message = "success"
			res.Data = data
		}

		c.JSON(200, res)
	}
}
