package main

import (
	"github.com/products-admin-service/internal/core/ports"
	"github.com/products-admin-service/internal/handlers"
	"github.com/products-admin-service/internal/server"
)

var (
	productsServer   ports.IServer
	productsHandlers ports.IHandlers
)

func init() {
	productsHandlers = handlers.NewHandlers()
	productsServer = server.NewServer(productsHandlers)
}

func main() {
	productsServer.Initialize()
}
