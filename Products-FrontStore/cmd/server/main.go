package main

import (
	pb "common/pkg/protos"
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/client"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/handlers"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/server"
)

var (
	productsFrontStoreServer     ports.IServer
	productsFrontStoreHandlers   ports.IProductsHandlers
	productsFrontStoreGrpcClient pb.ProductsServicesClient
	productsFrontStoreServices   ports.IProductsServices
)

func init() {
	ctx := context.TODO()
	productsFrontStoreGrpcClient = client.NewProductsGRPCClient()
	productsFrontStoreServices = services.NewProductsServices(ctx, productsFrontStoreGrpcClient)
	productsFrontStoreHandlers = handlers.NewHandlers(productsFrontStoreServices)
	productsFrontStoreServer = server.NewServer(productsFrontStoreHandlers)

}

func main() {
	productsFrontStoreServer.Initialize()
}
