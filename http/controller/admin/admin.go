package admin

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
