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
	"garavel/initialize"
	"github.com/google/wire"
	"net/http"
)

// init application.
func initializeApp(h *http.Server) (*App, error) {
	panic(
		wire.Build(
			initialize.ProviderSet,
			newApp,
			initServer,
		),
	)
}