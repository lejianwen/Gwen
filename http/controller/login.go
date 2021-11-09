package controller

import (
	"Gwen/global"
	"Gwen/model"
	"Gwen/model/request"
	"Gwen/model/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
	"time"
)

type Login struct {
}

// Login 登录
// @Tags 登录
// @Summary 登录
// @Description 登录
// @Accept  json
// @Produce  json
// @Param body body request.Login true "登录信息"
// @Success 200 {object} response.Response{data=model.Admin}
// @Failure 500 {object} response.Response
// @Router /login [post]
func (ct *Login) Login(c *gin.Context) {
	f := &request.Login{}
	err := c.ShouldBindJSON(f)
	if err != nil {
		global.Logger.Error(err)
		response.Fail(c, 101, "系统错误")
		return
	}

	errList := global.Validator.ValidStruct(f)
	if len(errList) > 0 {
		response.Fail(c, 101, errList[0])
		return
	}
	ad := service.AllService.AdminService.InfoByPwd(f.Username, f.Password)
	if ad.Id == 0 {
		response.Fail(c, 101, "用户名密码错误")
		return
	}
	if ad.Token == "" {
		ad.Token = service.AllService.AdminService.GenerateToken(ad)
		ad.TokenExpireTime = time.Now().Add(time.Duration(30*24*3600) * time.Second)
		global.DB.Save(ad)
	}

	response.Success(c, ad)
}

// Logout 登出
// @Tags 登录
// @Summary 登出
// @Description 登出
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /logout [post]
// @Security token
func (ct *Login) Logout(c *gin.Context) {
	a, ex := c.Get("curAdmin")
	if !ex {
		response.Fail(c, 101, "账号不存在")
	}
	if a == nil {
		response.Fail(c, 101, "账号不存在")
	}
	ad := a.(*model.Admin)
	ad.Token = ""
	global.DB.Save(ad)
	response.Success(c, nil)
}
