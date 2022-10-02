/*
@Time : 2022/10/2 18:19
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : router
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package initialize

import (
	"garavel/global"
	"garavel/global/middleware"
	"garavel/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(Routers)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// https
	//Router.Use(middleware.LoadTls())

	// 跨域
	Router.Use(middleware.Cors())

	RouterGroup := Router.Group("")
	// api
	routes.InitApiRouter(RouterGroup)

	// web
	Router.LoadHTMLGlob(global.G_CONFIG.System.ResourcePath)
	routes.InitWebRouter(RouterGroup)

	return Router
}
