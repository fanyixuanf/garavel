/*
@Time : 2022/9/25 17:49
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : api
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package router

import "github.com/gin-gonic/gin"

func InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	ApiRouter := Router.Group("api")
	{
		//ApiRouter.POST("register", v1.Register)
	}
	return ApiRouter
}
