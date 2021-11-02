package router

import (
	"Gwen/http/middleware"
	"Gwen/http/router/admin"
	"github.com/gin-gonic/gin"
)

func AdminInit(g *gin.Engine) {

	adg := g.Group("/admin")

	(&admin.Login{}).Init(adg)

	adg.Use(middleware.AdminAuth())
	(&admin.Admin{}).Init(adg)

}
