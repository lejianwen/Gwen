package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	AdminInit(g)

	ApiInit(g)
}
