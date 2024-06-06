package admin

import (
	"Gwen/http/controller/admin"
	"github.com/gin-gonic/gin"
)

type AdminRole struct {
	controller *admin.AdminRole
}

func (g *AdminRole) Bind(rg *gin.RouterGroup) {
	aR := rg.Group("/admin_role")
	{
		aR.GET("/detail/:id", g.controller.Detail)
		aR.GET("/list", g.controller.List)
		aR.POST("/create", g.controller.Create)
		aR.POST("/update", g.controller.Update)
		aR.POST("/delete", g.controller.Delete)
	}
}
