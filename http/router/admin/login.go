package admin

import (
	"Gwen/http/controller/admin"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
)

type Login struct {
	controller *admin.Login
}

func (g *Login) Bind(rg *gin.RouterGroup) {
	rg.POST("/login", g.controller.Login)
	rg.Use(middleware.AdminAuth()).POST("/logout", g.controller.Logout)
}
