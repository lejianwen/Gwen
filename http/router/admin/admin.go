package admin

import (
	"Gwen/http/controller"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	controller *controller.Admin
}

func (g *Admin) Bind(rg *gin.RouterGroup) {
	aR := rg.Group("/admin")
	{
		aR.GET("/list", g.controller.List)
		aR.GET("/detail/:id", g.controller.Detail)
		aR.POST("/create", g.controller.Create)
		aR.POST("/update", g.controller.Update)
		aR.POST("/delete", g.controller.Delete)
		aR.GET("/error", g.controller.Error)
	}
}
