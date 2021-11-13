package admin

import (
	"Gwen/http/controller"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
)

type Login struct {
	controller *controller.Login
}

func (g *Login) Init(rg *gin.RouterGroup) {
	rg.POST("/login", g.controller.Login)
	rg.Use(middleware.AdminAuth()).POST("/logout", g.controller.Logout)
}
