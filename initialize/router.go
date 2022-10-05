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
	"garavel/config"
	"garavel/global/middleware"
	"garavel/routes"
	"garavel/utils"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func Routers(config *config.Server, cors *middleware.Cors, routes *routes.Routers, tls *middleware.LoadTls) *gin.Engine {
	var Router = gin.Default()

	// https
	//Router.Use(tls.LoadTlsHandler())

	// 跨域
	Router.Use(cors.CorsHandler())

	RouterGroup := Router.Group("")
	// api
	routes.InitApiRouter(RouterGroup)

	// web
	//Router.LoadHTMLGlob(global.G_CONFIG.System.ResourcePath)
	Router.LoadHTMLGlob(filepath.Join(utils.RootPath(), "../", config.System.ResourcePath))
	Router.StaticFile("favicon.ico", "../resources/images/favicon.ico")
	routes.InitWebRouter(RouterGroup)

	return Router
}
