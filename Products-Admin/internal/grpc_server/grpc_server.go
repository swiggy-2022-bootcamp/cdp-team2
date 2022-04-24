package grpc_server

import (
	"log"
	"net"

	pb "common/protos/products"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/grpc_handlers"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpc_handlers.GRPCHandlers
	ProductsGrpcHandlers pb.ProductsServicesServer
}

// Check if GRPCServer implements ProductsServiceServer API
var _ ports.IGRPCServer = (*GRPCServer)(nil)
var _ pb.ProductsServicesServer = (*GRPCServer)(nil)

func NewGRPCServer(productsGrpcHandlers pb.ProductsServicesServer) *GRPCServer {
	return &GRPCServer{
		ProductsGrpcHandlers: productsGrpcHandlers,
	}
}

func (gs *GRPCServer) Initialize() {
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductsServicesServer(s, gs.ProductsGrpcHandlers)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
