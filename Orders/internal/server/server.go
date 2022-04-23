package server

import (
	"fmt"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/config"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Order/protos/order"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/services/grpc_services"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

// @title Checkout Api
// @version 1.0
// @description This is order service.

// @host localhost:7451
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Start() {
	router := InitRouter()

	endPoint := fmt.Sprintf(":%s", config.Server["PORT"])

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	log.Printf("start http server listening %s", endPoint)

	go startGrpcServer()

	_ = server.ListenAndServe()
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, &grpc_services.Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}