/*
@Time : 2022/10/2 18:02
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : app
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package bootstrap

import (
	"fmt"
	"garavel/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type App struct {
	config *config.Server
	log *zap.Logger
	httpServer *http.Server
}

func httpServer(config *config.Server, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", config.System.Addr),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func newApp(config *config.Server, log *zap.Logger, h *http.Server) *App {
	return &App {
		config: config,
		httpServer: h,
		log: log,
	}
}

func (app *App) Run() error {

	app.log.Info("port: ", zap.Any("port:", app.config.System.Addr))
	if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	return nil
}
