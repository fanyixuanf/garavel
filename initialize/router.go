/*
@Time : 2022/9/25 17:38
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
	"garavel/router"
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
	router.InitApiRouter(RouterGroup)

	// web
	Router.LoadHTMLGlob(global.G_CONFIG.System.ResourcePath)
	router.InitWebRouter(RouterGroup)

	return Router
}