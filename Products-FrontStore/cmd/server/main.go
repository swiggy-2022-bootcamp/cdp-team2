package main

import (
	"github.com/products-frontstore-service/internal/core/ports"
	"github.com/products-frontstore-service/internal/handlers"
	"github.com/products-frontstore-service/internal/server"
)

var (
	productsFrontStoreServer   ports.IServer
	productsFrontStoreHandlers ports.IHandlers
)

func init() {
	productsFrontStoreHandlers = handlers.NewHandlers()
	productsFrontStoreServer = server.NewServer(productsFrontStoreHandlers)
}

func main() {
	productsFrontStoreServer.Initialize()
}
