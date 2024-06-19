package admin

import (
	"Gwen/global"
	"Gwen/http/request"
	"Gwen/http/response"
	"Gwen/model"
	"Gwen/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdminRole struct {
}

// Detail 管理员角色
// @Tags 管理员角色
// @Summary 管理员角色详情
// @Description 管理员角色详情
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} response.Response{data=model.AdminRole}
// @Failure 500 {object} response.Response
// @Router /admin_role/detail/{id} [get]
// @Security token
func (ct *AdminRole) Detail(c *gin.Context) {
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)
	item := service.AllService.AdminRoleService.InfoById(uint(iid))
	if item.Id > 0 {
		response.Success(c, item)
		return
	}
	response.Fail(c, 101, "信息不存在")
	return
}

// Create 管理员角色创建
// @Tags 管理员角色
// @Summary 管理员角色创建
// @Description 管理员角色创建
// @Accept  json
// @Produce  json
// @Param body body request.AdminRoleForm true "管理员角色"
// @Success 200 {object} response.Response{data=model.AdminRole}
// @Failure 500 {object} response.Response
// @Router /admin_role/create [post]
// @Security token
func (ct *AdminRole) Create(c *gin.Context) {
	f := &request.AdminRoleForm{}

	if err := c.ShouldBindJSON(f); err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "系统错误")
		return
	}
	item := &model.AdminRole{
		SeeCb: 0,
		Name:  f.Name,
	}
	err := service.AllService.AdminRoleService.Create(item)
	if err != nil {
		response.Fail(c, 101, "创建失败")
		return
	}
	response.Success(c, item)
}

// List 列表
// @Tags 管理员角色
// @Summary 管理员角色列表
// @Description 管理员角色列表
// @Accept  json
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "页大小"
// @Success 200 {object} response.Response{data=model.AdminRoleListRes}
// @Failure 500 {object} response.Response
// @Router /admin_role/list [get]
// @Security token
func (ct *AdminRole) List(c *gin.Context) {
	query := &model.Pagination{}
	if err := c.ShouldBindQuery(query); err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "参数错误")
		return
	}
	res := service.AllService.AdminRoleService.List(uint(query.Page), uint(query.PageSize), nil)
	response.Success(c, res)
}

// Update 编辑
// @Tags 管理员角色
// @Summary 管理员角色编辑
// @Description 管理员角色编辑
// @Accept  json
// @Produce  json
// @Param body body request.AdminRoleForm true "管理员信息"
// @Success 200 {object} response.Response{data=model.AdminRole}
// @Failure 500 {object} response.Response
// @Router /admin_role/update [post]
// @Security token
func (ct *AdminRole) Update(c *gin.Context) {
	f := &request.AdminRoleForm{}
	if err := c.ShouldBindJSON(f); err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "系统错误")
		return
	}
	errList := global.Validator.ValidStruct(f)
	if len(errList) > 0 {
		response.Fail(c, 101, errList[0])
		return
	}
	item := service.AllService.AdminRoleService.InfoById(f.Id)
	if item.Id > 0 {
		v := map[string]interface{}{
			"name": f.Name,
		}
		err := service.AllService.AdminRoleService.Update(item, v)
		if err == nil {
			response.Success(c, nil)
			return
		}
		response.Fail(c, 101, err.Error())
		return
	}
	response.Fail(c, 101, "信息不存在")
}

// Delete 删除
// @Tags 管理员角色
// @Summary 管理员角色删除
// @Description 管理员角色删除
// @Accept  json
// @Produce  json
// @Param body body model.AdminRole true "管理员角色信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin_role/delete [post]
// @Security token
func (ct *AdminRole) Delete(c *gin.Context) {
	f := &model.AdminRole{}
	if err := c.ShouldBindJSON(f); err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "系统错误")
		return
	}
	id := f.Id
	errList := global.Validator.ValidVar(id, "required,gt=0")
	if len(errList) > 0 {
		response.Fail(c, 101, "参数错误")
		return
	}
	item := service.AllService.AdminRoleService.InfoById(f.Id)
	if item.Id > 0 {
		err := service.AllService.AdminRoleService.Delete(item)
		if err == nil {
			response.Success(c, nil)
			return
		}
		response.Fail(c, 101, err.Error())
		return
	}
	response.Fail(c, 101, "信息不存在")
}
