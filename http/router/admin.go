package router

import (
	_ "Gwen/docs/admin"
	"Gwen/global"
	"Gwen/http/controller/admin"
	"Gwen/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Init(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})
	//swagger
	//g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName("admin")))

	adg := g.Group("/admin-api")
	LoginBind(adg)

	adg.Use(middleware.AdminAuth())
	FileBind(adg)
	AdminBind(adg)
	AdminRoleBind(adg)

	//访问静态文件
	g.StaticFS("/upload", http.Dir(global.Config.Gin.ResourcesPath+"/upload"))
}
func LoginBind(rg *gin.RouterGroup) {
	cont := &admin.Login{}
	rg.POST("/login", cont.Login)
	rg.POST("/logout", cont.Logout)
}

func AdminBind(rg *gin.RouterGroup) {
	aR := rg.Group("/admin")
	{
		cont := &admin.Admin{}
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
		cont := &admin.AdminRole{}
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
		cont := &admin.File{}
		aR.POST("/notify", cont.Notify)
		aR.OPTIONS("/oss_token", nil)
		aR.OPTIONS("/upload", nil)
		aR.GET("/oss_token", cont.OssToken)
		aR.POST("/upload", cont.Upload)
	}
}
