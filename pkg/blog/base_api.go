package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthBaseAPI(c *gin.Context) error {
	token := c.GetHeader("token")
	success, err := AuthService(token)
	if err != nil {
		return err
	}
	if !success {
		return fmt.Errorf("auth failed")
	}
	return nil
}
