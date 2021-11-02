//+build !windows

package http

import (
	"Gwen/global"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func Run(g *gin.Engine) {
	endless.ListenAndServe(global.Config.Gin.Addr, g)
}
