package http

import (
	"Gwen/global"
	"Gwen/http/middleware"
	"Gwen/http/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const UPLOAD_PATH = "/upload/"

func Init() {
	gin.SetMode(global.Config.Gin.Mode)
	g := gin.New()

	if global.Config.Gin.Mode == gin.ReleaseMode {
		//修改gin Recovery日志 输出为logger的输出点
		if global.Logger != nil {
			gin.DefaultErrorWriter = global.Logger.WriterLevel(logrus.ErrorLevel)
		}
	}

	g.Use(middleware.Logger(), gin.Recovery())

	router.Init(g)
	Run(g)
}

func ApiInit() {
	gin.SetMode(global.Config.Gin.Mode)
	g := gin.New()

	if global.Config.Gin.Mode == gin.ReleaseMode {
		//修改gin Recovery日志 输出为logger的输出点
		if global.Logger != nil {
			gin.DefaultErrorWriter = global.Logger.WriterLevel(logrus.ErrorLevel)
		}
	}

	g.Use(middleware.Logger(), gin.Recovery())

	router.ApiInit(g)
	Run(g)
}
