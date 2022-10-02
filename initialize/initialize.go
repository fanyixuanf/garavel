/*
@Time : 2022/10/2 21:06
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : initialize
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package initialize

import "github.com/google/wire"

var ProviderSet = wire.NewSet(Routers)
