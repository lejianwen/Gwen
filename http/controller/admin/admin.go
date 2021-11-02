package admin

import (
	"Gwen/global"
	"Gwen/http/controller"
	"Gwen/model"
	"Gwen/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	service *service.AdminService
}

func (ct *Admin) Index(c *gin.Context) {
	admin := service.AllService.AdminService.Info(1)
	global.Cache.Set("ted", "ddd", 300)
	controller.Success(c, admin)
}

func (ct *Admin) Create(c *gin.Context) {
	admin := &model.Admin{
		Nickname: "test",
		Username: "go-test",
	}
	controller.Success(c, admin)
}

func (ct *Admin) List(c *gin.Context) {
	list := service.AllService.AdminService.List()
	controller.ListResponse(c, list, 1, 100)
}

func (ct *Admin) Error(c *gin.Context) {
	fmt.Println(ct.service)

}
