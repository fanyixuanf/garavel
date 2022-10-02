//go:build wireinject
// +build wireinject

/*
@Time : 2022/10/1 18:49
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : wire.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package boostrap

import (
	appConfig "garavel/config"
	"garavel/initialize"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// init application.
func InitializeApp(*appConfig.System, *zap.Logger) (*App, error) {
	panic(
		wire.Build(
			initialize.ProviderSet,
			newApp,
			initServer,
		),
	)
}