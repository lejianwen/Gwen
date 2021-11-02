package http

import (
	"Gwen/global"
	"Gwen/http/middleware"
	"Gwen/http/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() {
	gin.SetMode(global.Config.Gin.Mode)
	g := gin.New()

	//修改gin Recovery日志 输出为logger的输出点
	if global.LOGGER != nil {
		gin.DefaultErrorWriter = global.LOGGER.WriterLevel(logrus.ErrorLevel)
	}

	g.Use(middleware.Logger(), gin.Recovery())

	router.Init(g)
	Run(g)
}
