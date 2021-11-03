package admin

import "github.com/gin-gonic/gin"

type Login struct {
}

// @Summary 登录
// @Description 登录
// @Accept  json
// @Produce  json
// @Param  body body controller.Response true "登录信息"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /admin/login [post]
//Login 登录
func (ct *Login) Login(c *gin.Context) {

}

func (ct *Login) Logout(c *gin.Context) {

}
