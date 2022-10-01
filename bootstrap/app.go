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
	appConfig "garavel/config"
	"github.com/gin-gonic/gin"
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
	go func() {
		if err := app.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}


type server interface {
	ListenAndServe() error
}

type App struct {
	config *appConfig.System
	httpSrv *http.Server
}

func newApp(h *http.Server) *App {
	return &App{
		httpSrv:h,
	}
}

func initServer(router *gin.Engine) server {
	return &http.Server{
		Addr:           ":8088",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}