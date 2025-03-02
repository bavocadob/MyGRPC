package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (r *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := getAuthToken(c)
		if t == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			if _, err := r.gRPCClient.VerifyAuth(t); err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
			} else {
				c.JSON(http.StatusOK, nil)
				c.Next()
			}
		}
	}
}

func getAuthToken(c *gin.Context) string {
	var token string
	authToken := c.Request.Header.Get("Authorization")
	authSlice := strings.Split(authToken, " ")
	if len(authSlice) != 2 {

	} else {
		token = authSlice[1]
	}

	return token
}
