package admin

import (
	"Gwen/http/controller/admin"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	controller *admin.Admin
}

func (g *Admin) Init(rg *gin.RouterGroup) {
	aR := rg.Group("/admin")
	{
		aR.GET("/list", g.controller.List)
		aR.GET("/index", g.controller.Index)
		aR.POST("/create", g.controller.Create)
		aR.GET("/error", g.controller.Error)
	}
}
