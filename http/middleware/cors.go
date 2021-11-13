package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "api-token")
		c.Header("Access-Control-Allow-Methods", c.Request.Method)
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
