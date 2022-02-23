package controller

import (
	"Gwen/global"
	"Gwen/http/request"
	"Gwen/http/response"
	"Gwen/model"
	"Gwen/service"
	"Gwen/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	admin := service.AllService.AdminService.Info(uint(iid))
	if admin.Id > 0 {
		admin.Token = "" // 去掉token
		response.Success(c, admin)
		return
	}
	response.Fail(c, 101, "信息不存在")
	return
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
		Status:   f.Status,
	}
	err := service.AllService.AdminService.Create(admin)
	if err != nil {
		response.Fail(c, 101, "创建失败")
		return
	}
	response.Success(c, admin)
}

//List 列表
// @Tags 管理员
// @Summary 管理员列表
// @Description 管理员列表
// @Accept  json
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "页大小"
// @Param nickname query int false "昵称"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/list [get]
// @Security token
func (ct *Admin) List(c *gin.Context) {
	query := &request.AdminListQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "参数错误")
		return
	}
	res := service.AllService.AdminService.List(query.Page, query.PageSize, func(tx *gorm.DB) {
		if query.Nickname != "" {
			tx.Where("nickname like ?", "%"+query.Nickname+"%")
		}
	})
	response.Success(c, res)
}

//Update 编辑
// @Tags 管理员
// @Summary 管理员编辑
// @Description 管理员编辑
// @Accept  json
// @Produce  json
// @Param body body request.AdminForm true "管理员信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/update [post]
// @Security token
func (ct *Admin) Update(c *gin.Context) {
	f := &request.AdminForm{}
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
	admin := service.AllService.AdminService.Info(f.Id)
	if admin.Id > 0 {
		v := map[string]interface{}{
			"role_id":  f.RoleId,
			"username": f.Username,
			"nickname": f.Nickname,
			"status":   f.Status,
		}
		err := service.AllService.AdminService.Update(admin, v)
		if err == nil {
			response.Success(c, nil)
			return
		}
		response.Fail(c, 101, err.Error())
		return
	}
	response.Fail(c, 101, "信息不存在")
}

//Delete 删除
// @Tags 管理员
// @Summary 管理员删除
// @Description 管理员编删除
// @Accept  json
// @Produce  json
// @Param body body request.AdminForm true "管理员信息"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/delete [post]
// @Security token
func (ct *Admin) Delete(c *gin.Context) {
	f := &request.AdminForm{}
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
	admin := service.AllService.AdminService.Info(f.Id)
	if admin.Id > 0 {
		err := service.AllService.AdminService.Delete(admin)
		if err == nil {
			response.Success(c, nil)
			return
		}
		response.Fail(c, 101, err.Error())
		return
	}
	response.Fail(c, 101, "信息不存在")
}

func (ct *Admin) Error(c *gin.Context) {
	panic("测试")
}
