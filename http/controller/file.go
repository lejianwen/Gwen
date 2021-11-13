package controller

import (
	"Gwen/global"
	"Gwen/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type File struct {
}

//OssToken 文件
// @Tags 文件
// @Summary 获取ossToken
// @Description 获取ossToken
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /file/oss_token [get]
// @Security token
func (f *File) OssToken(c *gin.Context) {
	token := global.Oss.GetPolicyToken("")
	response.Success(c, token)
}

func (f *File) Notify(c *gin.Context) {
	res := global.Oss.Verify(c.Request)
	if !res {
		response.Fail(c, 101, "权限错误")
		return
	}
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	// todo 入库等
	for k, v := range c.Request.PostForm {

		fmt.Printf("k:%v\n", k)

		fmt.Printf("v:%v\n", v)

	}

	response.Success(c, c.Request.PostForm)

}
