package http

import (
	"Gwen/global"
	"Gwen/http/middleware"
	"Gwen/http/router"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin.SetMode(global.Config.Gin.Mode)
	g := gin.New()

	//修改gin Recovery日志 输出为logger的输出点
	if global.Logger != nil {
		//gin.DefaultErrorWriter = global.Logger.WriterLevel(logrus.ErrorLevel)
	}

	g.Use(gin.Logger())

	g.Use(middleware.Logger(), gin.Recovery())

	router.Init(g)
	Run(g)
}
