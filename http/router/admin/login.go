package admin

import (
	"Gwen/http/controller/admin"
	"github.com/gin-gonic/gin"
)

type Login struct {
	controller *admin.Login
}

func (g *Login) Init(adg *gin.RouterGroup) {
	adg.POST("/login", g.controller.Login)
	adg.POST("/logout", g.controller.Logout)
}
