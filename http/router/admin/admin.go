package admin

import (
	controller2 "Gwen/http/controller"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	controller *controller2.Admin
}

func (g *Admin) Init(rg *gin.RouterGroup) {
	aR := rg.Group("/admin")
	{
		aR.GET("/list", g.controller.List)
		aR.GET("/detail/:id", g.controller.Detail)
		aR.POST("/create", g.controller.Create)
		aR.GET("/error", g.controller.Error)
	}
}
