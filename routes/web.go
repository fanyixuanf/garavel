/*
@Time : 2022/10/2 18:20
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : web
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (r *Routers) InitWebRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	webRouter := router.Group("")
	{
		webRouter.GET("/", func(c *gin.Context) {
			c.HTML(200, "view/welcome.html", gin.H{"title": "Garavel", "address": fmt.Sprintf("http://127.0.0.1:%d", 8088)})
		})
	}
	return webRouter
}