package admin

import (
	"Gwen/http/controller"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
)

type Login struct {
	controller *controller.Login
}

func (g *Login) Init(adg *gin.RouterGroup) {
	adg.POST("/login", g.controller.Login)
	adg.Use(middleware.AdminAuth()).POST("/logout", g.controller.Logout)
}
