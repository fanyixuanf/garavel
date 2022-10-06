/*
@Time : 2022/10/2 18:20
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : api
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package routes

import (
	"github.com/gin-gonic/gin"
)

func (r *Routers)InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	ApiRouter := Router.Group("api")
	{
		//ApiRouter.GET("/register", web.Welcome)
	}
	return ApiRouter
}
