package admin

import (
	controller2 "Gwen/http/controller"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
)

type File struct {
	controller *controller2.File
}

func (g *File) Init(rg *gin.RouterGroup) {
	aR := rg.Group("/file").Use(middleware.Cors())
	{
		aR.POST("/notify", g.controller.Notify)
		aR.OPTIONS("/oss_token", nil)
		aR.Use(middleware.AdminAuth()).GET("/oss_token", g.controller.OssToken)
	}
}
