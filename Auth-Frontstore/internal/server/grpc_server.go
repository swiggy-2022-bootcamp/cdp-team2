package server

import (
	"log"
	"net"

	"github.com/auth-frontstore-service/config"
	"github.com/auth-frontstore-service/internal/handlers"
	"github.com/auth-frontstore-service/protos/authpb"
	"google.golang.org/grpc"
)

func (server *Server) InitializeGrpcServer() {

	lis, err := net.Listen("tcp", "0.0.0.0:"+config.Config["GRPC"])
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &handlers.GrpcServer{
		authpb.UnimplementedAuthServiceServer{},
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
