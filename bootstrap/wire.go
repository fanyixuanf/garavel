//go:build wireinject
// +build wireinject

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
	"garavel/global/middleware"
	"garavel/initialize"
	"garavel/routes"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitializeApp(*config.Server, *zap.Logger) (*App,func(), error) {
	panic(
		wire.Build(
			initialize.ProviderSet,
			middleware.ProviderSet,
			routes.ProviderSet,
			newApp,
			httpServer,
		),
	)
}