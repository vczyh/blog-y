package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginAPI(c *gin.Context) (interface{}, error) {
	var req = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&req); err != nil {
		return nil, fmt.Errorf("param error")
	}
	return LoginService(req.Username, req.Password)
}

func AuthAPI(c *gin.Context) (interface{}, error) {
	token := c.GetHeader("token")
	return AuthService(token)
}
