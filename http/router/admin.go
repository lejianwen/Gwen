package router

import (
	"Gwen/global"
	"Gwen/http/middleware"
	"Gwen/http/router/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminInit(g *gin.Engine) {

	(&admin.Login{}).Init(g.Group("/admin-api"))
	(&admin.File{}).Init(g.Group("/admin-api"))

	adg := g.Group("/admin-api")
	adg.Use(middleware.AdminAuth())
	(&admin.Admin{}).Init(adg)
	(&admin.AdminRole{}).Init(adg)

	g.StaticFS("/upload", http.Dir(global.Config.Gin.ResourcesPath+"/upload"))
}
