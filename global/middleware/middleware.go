/*
@Time : 2022/10/3 18:42
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : middleware
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package middleware

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewCors)
