package main

import (
	"github.com/auth-frontstore-service/internal/core/ports"
	"github.com/auth-frontstore-service/internal/handlers"
	"github.com/auth-frontstore-service/internal/middlewares"
	"github.com/auth-frontstore-service/internal/server"
)

var (
	frontStoreAuthServer      ports.IServer
	frontStoreAuthHandlers    ports.IHandlers
	frontStoreAuthMiddlewares ports.IMiddlewares
)

func init() {
	frontStoreAuthHandlers = handlers.NewHandlers()
	frontStoreAuthMiddlewares = middlewares.NewMiddlewares()
	frontStoreAuthServer = server.NewServer(frontStoreAuthHandlers, frontStoreAuthMiddlewares)
}

func main() {
	go frontStoreAuthServer.InitializeGrpcServer()
	frontStoreAuthServer.Initialize()
}
