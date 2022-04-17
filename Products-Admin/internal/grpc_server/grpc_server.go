package grpc_server

import (
	"log"
	"net"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/grpc_handlers"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/protos"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedProductsServicesServer
	grpc_handlers.GRPCHandlers
}

var _ ports.IGRPCServer = (*GRPCServer)(nil)

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (gs *GRPCServer) Initialize() {
	lis, err := net.Listen("tcp", ":7459")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductsServicesServer(s, *gs)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
