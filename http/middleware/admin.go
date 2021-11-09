package middleware

import (
	"Gwen/model"
	"Gwen/model/response"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

// AdminAuth 后台权限验证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("api-token")
		if token == "" {
			response.Fail(c, 101, "请先登录")
			c.Abort()
			return
		}
		admin, res := service.AllService.AdminService.CheckToken(token)
		if !res {
			response.Fail(c, 101, "请先登录")
			c.Abort()
			return
		}
		if admin.Status != model.COMMON_STATUS_ENABLE {
			response.Fail(c, 101, "账号已被禁用")
			c.Abort()
			return
		}
		c.Set("curAdmin", admin)
		c.Next()
	}
}
