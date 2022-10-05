// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"garavel/config"
	"garavel/global/middleware"
	"garavel/initialize"
	"garavel/routes"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InitializeApp(server *config.Server, logger *zap.Logger) (*App, func(), error) {
	cors := middleware.NewCors()
	routers := routes.NewRouter()
	loadTls := middleware.NewLoadTls()
	engine := initialize.Routers(server, cors, routers, loadTls)
	server2 := httpServer(server, engine)
	app := newApp(server, logger, server2)
	return app, func() {
	}, nil
}
