package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/handlers"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/server"
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
