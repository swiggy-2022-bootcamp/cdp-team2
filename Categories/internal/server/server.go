package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "common/protos/category"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/config"
	catGrpc "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @title Categories Api
// @version 1.0
// @description This is categories crud service.

// @host localhost:7450
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Start() {
	router, err := InitRouter()

	if err != nil {
		log.Println("Error starting server")
		log.Println(err.Error())
		return
	}

	endPoint := fmt.Sprintf(":%s", config.Server["PORT"])

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	log.Printf("start http server listening %s", endPoint)

	go StartGrpcServer()

	_ = server.ListenAndServe()
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":7459")
	if err != nil {
		log.Fatalf("GRPC failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cats, err := catGrpc.NewCategoryGrpcServer()
	if err != nil {
		log.Fatalf("GRPC failed to start grpc service: %v", err)
		return
	}
	pb.RegisterCategoryServiceServer(s, cats)
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("GRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("GRPC failed to serve: %v", err)
	}
}
