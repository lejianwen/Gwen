package router

import (
	_ "Gwen/docs/api"
	"Gwen/global"
	"Gwen/http/controller/api"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func ApiInit(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})
	//swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName("api")))

	frg := g.Group("/api")
	i := &api.Index{}
	frg.GET("/", i.Index)

	u := &api.User{}
	frg.Use(middleware.JwtAuth())
	frg.GET("/user/info", u.Info)

	//访问静态文件
	g.StaticFS("/upload", http.Dir(global.Config.Gin.ResourcesPath+"/upload"))
}
