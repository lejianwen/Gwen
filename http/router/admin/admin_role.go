package admin

import (
	controller2 "Gwen/http/controller"
	"github.com/gin-gonic/gin"
)

type AdminRole struct {
	controller *controller2.AdminRole
}

func (g *AdminRole) Init(rg *gin.RouterGroup) {
	aR := rg.Group("/admin_role")
	{
		aR.GET("/detail/:id", g.controller.Detail)
		aR.GET("/list", g.controller.List)
		aR.POST("/create", g.controller.Create)
		aR.POST("/update", g.controller.Update)
		aR.POST("/delete", g.controller.Delete)
	}
}
