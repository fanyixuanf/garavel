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
	"garavel/global"
	"garavel/initialize"
	"time"
)

func Run() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.G_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	fmt.Println(s.ListenAndServe().Error())
}


type server interface {
	ListenAndServe() error
}

