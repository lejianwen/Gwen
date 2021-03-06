package middleware

import (
	"Gwen/http/response"
	"Gwen/model"
	"Gwen/service"
	"github.com/gin-gonic/gin"
)

// AdminAuth 后台权限验证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		//测试先关闭
		token := c.GetHeader("api-token")
		if token == "" {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}
		admin, res := service.AllService.AdminService.CheckToken(token)
		if !res {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}
		if admin.Status != model.COMMON_STATUS_ENABLE {
			response.Fail(c, 403, "请先登录")
			c.Abort()
			return
		}

		c.Set("curAdmin", admin)

		c.Next()
	}
}
