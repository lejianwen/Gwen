package api

import (
	"Gwen/http/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

type User struct {
}

// Info 用户信息
// @Tags 用户
// @Summary 用户信息
// @Description 用户信息
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=response.UserInfoData}
// @Failure 500 {object} response.Response
// @Router /user/info [get]
// @Security token
func (u *User) Info(c *gin.Context) {
	user := service.AllService.UserService.CurUser(c)
	response.Success(c, &response.UserInfoData{
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
	})
}
