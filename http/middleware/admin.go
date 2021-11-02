package middleware

import (
	"Gwen/http/controller"
	"github.com/gin-gonic/gin"
)

// AdminAuth 后台权限验证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//简单测试
		key := c.Query("key")
		if key == "" {
			controller.Fail(c, 101, "错误")
			c.Abort()
			return
		}
		c.Next()
	}
}
