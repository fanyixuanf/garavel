

/*
@Time : 2022/10/2 18:24
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : wrie
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package bootstrap

import (
	"garavel/config"
	"garavel/initialize"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitializeApp(*config.Server, *zap.Logger) (*App, error) {
	panic(
		wire.Build(
			initialize.ProviderSet,
			newApp,
			httpServer,
		),
	)
}