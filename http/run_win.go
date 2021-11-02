//+build windows

package http

import (
	"Gwen/global"
	"github.com/gin-gonic/gin"
)

func Run(g *gin.Engine) {
	g.Run(global.Config.Gin.Addr)
}
