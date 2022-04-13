package main

import (
	"github.com/auth-admin-service/internal/core/ports"
	"github.com/auth-admin-service/internal/handlers"
	"github.com/auth-admin-service/internal/server"
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
