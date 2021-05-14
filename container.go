//+build wireinject

package main

import (
	"github.com/demogoo/wirego/internal/config"
	"github.com/demogoo/wirego/internal/handlers"
	"github.com/demogoo/wirego/internal/server"
	"github.com/demogoo/wirego/internal/services"
	"github.com/demogoo/wirego/internal/storage"
	"github.com/google/wire"
)

func CreateServer() *server.Server {
	panic(wire.Build(
		config.NewConfig,
		storage.NewDBConn,
		storage.NewCacheConn,
		storage.NewMongoDBConn,
		services.NewItemService,
		services.NewOrderService,
		services.NewMemberService,
		handlers.NewItemHandler,
		handlers.NewOrderHandler,
		handlers.NewMemberHandler,
		server.NewServer,
	))
}
