/*
@Time : 2022/9/25 16:38
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : app.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package boostrap

import (
	"fmt"
	appConfig "garavel/config"
	"garavel/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)


func (app *App) Run() {

	//// initialize server
	//
	//Router := initialize.Routers()
	//address := fmt.Sprintf(":%d", global.G_CONFIG.System.Addr)
	//s := initServer(address, Router)
	//time.Sleep(10 * time.Microsecond)
	//fmt.Println(s.ListenAndServe().Error())



	if err := app.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}


	//app.logger.Info("success")
	//go func() {
	//	if err := app.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		panic(err)
	//	}
	//	app.logger.Info("aaaa")
	//}()
}

type App struct {
	config *appConfig.System
	logger *zap.Logger
	httpSrv *http.Server
}

func newApp(c *appConfig.System, logger *zap.Logger, h *http.Server) *App {
	return &App{
		config: c,
		logger: logger,
		httpSrv: h,
	}
}

func initServer(c *appConfig.System, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", global.G_CONFIG.System.Addr),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}