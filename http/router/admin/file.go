package admin

import (
	"Gwen/http/controller"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
)

type File struct {
	controller *controller.File
}

func (g *File) Bind(rg *gin.RouterGroup) {
	aR := rg.Group("/file").Use(middleware.Cors())
	{
		aR.POST("/notify", g.controller.Notify)
		aR.OPTIONS("/oss_token", nil)
		aR.OPTIONS("/upload", nil)
		aR.Use(middleware.AdminAuth()).GET("/oss_token", g.controller.OssToken)
		aR.POST("/upload", g.controller.Upload)
	}
}
