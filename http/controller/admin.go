package controller

import (
	"Gwen/global"
	"Gwen/model"
	"Gwen/model/request"
	"Gwen/model/response"
	"Gwen/service"
	"Gwen/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Admin struct {
}

//Detail 管理员
// @Tags 管理员
// @Summary 管理员详情
// @Description 管理员详情
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/detail/{id} [get]
// @Security token
func (ct *Admin) Detail(c *gin.Context) {
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)
	admin := service.AllService.AdminService.Info(iid)
	admin.Token = "" // 去掉token
	response.Success(c, admin)
}

//Create 管理员
// @Tags 管理员
// @Summary 创建管理员
// @Description 创建管理员
// @Accept  json
// @Produce  json
// @Param body body request.AdminForm true "管理员信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/create [post]
// @Security token
func (ct *Admin) Create(c *gin.Context) {
	f := &request.AdminForm{}

	if err := c.ShouldBindJSON(f); err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "系统错误")
		return
	}
	admin := &model.Admin{
		RoleId:   f.RoleId,
		Username: f.Username,
		Password: utils.Md5(f.Password),
		Nickname: f.Nickname,
		Status:   1,
	}
	res := global.DB.Create(admin)
	if res.Error != nil {
		response.Fail(c, 101, "创建失败")
		return
	}
	response.Success(c, admin)
}

//List 登录
// @Tags 管理员
// @Summary 管理员列表
// @Description 管理员列表
// @Accept  json
// @Produce  json
// @Param page_num query int false "页码"
// @Param page_size query int false "页大小"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/list [get]
// @Security token
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
