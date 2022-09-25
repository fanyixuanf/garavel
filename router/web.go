/*
@Time : 2022/9/25 17:53
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : web
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package router

import (
	"garavel/app/Http/Controllers/web"
	"github.com/gin-gonic/gin"
)

func InitWebRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	webRouter := router.Group("")
	{
		webRouter.GET("/", web.Welcome)
	}
	return webRouter
}
