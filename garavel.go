/*
@Time : 2022/9/25 16:00
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : garavel.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package main

import (
	"garavel/bootstrap"
	"garavel/config"
	"go.uber.org/zap"
)

func main () {
	//initializeApp()

	//msg := NewMessage()
	//greeter := NewGreeter(msg)
	//evet := NewEvent(greeter)
	//evet.start()
	app, err := boostrap.InitializeApp(conf, zapa)
	if err != nil {
		zapa.Info("New App error")
	}
	app.Run()
}

var (
	conf *config.System
	zapa *zap.Logger
)
