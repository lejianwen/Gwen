package admin

import (
	"Gwen/model"
	"Gwen/model/request"
	"Gwen/model/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

type Login struct {
}

// @Summary 登录
// @Description 登录
// @Accept  json
// @Produce  json
// @Param  body body request.Login true "登录信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/login [post]
//Login 登录
func (ct *Login) Login(c *gin.Context) {
	f := &request.Login{}
	c.ShouldBindJSON(f)
	ad := service.AllService.AdminService.Info(1)
	response.Success(c, &struct {
		Token string       `json:"token"`
		Admin *model.Admin `json:"admin,omitempty"`
	}{
		Token: "test",
		Admin: ad,
	})
}

func (ct *Login) Logout(c *gin.Context) {

}
