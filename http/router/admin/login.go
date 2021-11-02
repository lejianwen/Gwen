package admin

import (
	"Gwen/http/controller/admin"
	"github.com/gin-gonic/gin"
)

type Login struct {
	controller *admin.Login
}

func (g *Login) Init(rg *gin.RouterGroup) {
	rg.POST("/login", g.controller.Login)
	rg.POST("/logout", g.controller.Logout)
}
