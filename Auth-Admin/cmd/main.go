package main

import (
	"github.com/auth-admin-service/internal/core/ports"
	"github.com/auth-admin-service/internal/handlers"
	"github.com/auth-admin-service/internal/middlewares"
	"github.com/auth-admin-service/internal/server"
)

var (
	adminAuthServer      ports.IServer
	adminAuthHandlers    ports.IHandlers
	adminAuthMiddlewares ports.IMiddlewares
)

func init() {
	adminAuthHandlers = handlers.NewHandlers()
	adminAuthMiddlewares = middlewares.NewMiddlewares()
	adminAuthServer = server.NewServer(adminAuthHandlers, adminAuthMiddlewares)
}

func main() {
	adminAuthServer.Initialize()
}
