package admin

import (
	controller2 "Gwen/http/controller"
	"github.com/gin-gonic/gin"
)

type Login struct {
	controller *controller2.Login
}

func (g *Login) Init(adg *gin.RouterGroup) {
	adg.POST("/login", g.controller.Login)
	adg.POST("/logout", g.controller.Logout)
}
