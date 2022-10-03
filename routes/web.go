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
	"garavel/app/Http/Controllers/web"
	"github.com/gin-gonic/gin"
)

func (r *Routers) InitWebRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	webRouter := router.Group("")
	{
		webRouter.GET("/", web.Welcome)
	}
	return webRouter
}