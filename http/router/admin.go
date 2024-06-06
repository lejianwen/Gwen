package router

import (
	"Gwen/global"
	"Gwen/http/middleware"
	"Gwen/http/router/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminInit(g *gin.Engine) {

	adg := g.Group("/admin-api")
	(&admin.Login{}).Bind(adg)

	adg.Use(middleware.AdminAuth())
	(&admin.File{}).Bind(adg)
	(&admin.Admin{}).Bind(adg)
	(&admin.AdminRole{}).Bind(adg)

	//访问静态文件
	g.StaticFS("/upload", http.Dir(global.Config.Gin.ResourcesPath+"/upload"))
}
