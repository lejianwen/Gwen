package router

import (
	"Gwen/http/controller"
	"github.com/gin-gonic/gin"
)

func LoginBind(rg *gin.RouterGroup) {
	cont := &controller.Login{}
	rg.POST("/login", cont.Login)
	rg.POST("/logout", cont.Logout)
}

func AdminBind(rg *gin.RouterGroup) {
	aR := rg.Group("/admin")
	{
		cont := &controller.Admin{}
		aR.GET("/list", cont.List)
		aR.GET("/detail/:id", cont.Detail)
		aR.POST("/create", cont.Create)
		aR.POST("/update", cont.Update)
		aR.POST("/delete", cont.Delete)
		aR.GET("/error", cont.Error)
	}
}

func AdminRoleBind(rg *gin.RouterGroup) {
	aR := rg.Group("/admin_role")
	{
		cont := &controller.AdminRole{}
		aR.GET("/list", cont.List)
		aR.GET("/detail/:id", cont.Detail)
		aR.POST("/create", cont.Create)
		aR.POST("/update", cont.Update)
		aR.POST("/delete", cont.Delete)
	}
}

func FileBind(rg *gin.RouterGroup) {
	aR := rg.Group("/file")
	{
		cont := &controller.File{}
		aR.POST("/notify", cont.Notify)
		aR.OPTIONS("/oss_token", nil)
		aR.OPTIONS("/upload", nil)
		aR.GET("/oss_token", cont.OssToken)
		aR.POST("/upload", cont.Upload)
	}
}
