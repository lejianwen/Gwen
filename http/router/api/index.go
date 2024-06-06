package api

import (
	"Gwen/http/controller/api"
	"github.com/gin-gonic/gin"
)

type Index struct {
	controller *api.Index
}

func (g *Index) Bind(rg *gin.RouterGroup) {
	aR := rg.Group("/index")
	{
		aR.GET("/", g.controller.Index)
	}
}
