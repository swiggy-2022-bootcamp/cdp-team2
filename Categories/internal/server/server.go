package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "common/protos/category"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/config"
	"google.golang.org/grpc"
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
	router := InitRouter()

	endPoint := fmt.Sprintf(":%s", config.Server["PORT"])

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	log.Printf("start http server listening %s", endPoint)

	go StartGrpcServer()

	_ = server.ListenAndServe()
}

type GrpcServer struct {
	pb.UnimplementedCategoryServiceServer
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":7459")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCategoryServiceServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
