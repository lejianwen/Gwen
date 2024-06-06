package api

import (
	"Gwen/http/response"
	"github.com/gin-gonic/gin"
)

type Index struct {
}

// Index Index
// @Tags Index
// @Summary Index
// @Description Index
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /index [get]
func (i *Index) Index(c *gin.Context) {
	response.Success(c, nil)
}
