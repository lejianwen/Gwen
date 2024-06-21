package middleware

import (
	"Gwen/global"
	"Gwen/http/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//测试先关闭
		token := c.GetHeader("api-token")
		if token == "" {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}
		uid, err := global.Jwt.ParseToken(token)
		if err != nil {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}
		if uid == 0 {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}

		user := service.AllService.UserService.InfoById(uid)
		//user := &model.User{
		//	Id:       uid,
		//	Nickname: "测试用户",
		//}
		if user.Id == 0 {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}
		if !service.AllService.UserService.CheckUserEnable(user) {
			response.Fail(c, 101, "你已被禁用")
			c.Abort()
			return
		}
		c.Set("curUser", user)

		c.Next()
	}
}
