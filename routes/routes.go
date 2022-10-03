/*
@Time : 2022/10/3 19:00
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : routes
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package routes

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewRouter)

type Routers struct {}

func NewRouter() *Routers {
	return &Routers{}
}