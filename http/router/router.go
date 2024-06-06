package router

import (
	_ "Gwen/docs"
	"Gwen/global"
	"Gwen/http/middleware"
	"Gwen/http/router/admin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Init(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})
	//swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adg := g.Group("/admin-api")
	(&admin.Login{}).Bind(adg)

	adg.Use(middleware.AdminAuth())
	(&admin.File{}).Bind(adg)
	(&admin.Admin{}).Bind(adg)
	(&admin.AdminRole{}).Bind(adg)

	//访问静态文件
	g.StaticFS("/upload", http.Dir(global.Config.Gin.ResourcesPath+"/upload"))
}
