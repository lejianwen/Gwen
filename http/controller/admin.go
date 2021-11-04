package controller

import (
	"Gwen/global"
	"Gwen/model"
	"Gwen/model/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (ct *Admin) Index(c *gin.Context) {
	admin := service.AllService.AdminService.Info(1)
	global.Cache.Set("ted", "ddd", 300)
	response.Success(c, admin)
}

func (ct *Admin) Create(c *gin.Context) {
	admin := &model.Admin{
		Nickname: "test",
		Username: "go-test",
	}
	response.Success(c, admin)
}

// @Summary 管理员列表
// @Description 管理员列表
// @Accept  json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "页大小"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/admin/list [get]
//List 登录
func (ct *Admin) List(c *gin.Context) {
	res := service.AllService.AdminService.List()
	//for _, admin := range res.Admins {
	//	admin.Role = nil
	//}
	response.Success(c, res)
}

func (ct *Admin) Error(c *gin.Context) {
	panic("测试")
}
