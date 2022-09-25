/*
@Time : 2022/9/25 17:59
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : server
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package boostrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
