package server

import (
	"fmt"
	"net"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"google.golang.org/grpc"

	"log"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/docs"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/protos/shipping"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/services/grpc_server_services"
)

// @title Shipping Address Microservice
// @version 1.0
// @description This is shipping address service.

// @host localhost:8003
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

type server struct {
	pb.UnimplementedServiceServer
}

func Start() {
	router := InitRouter()

	docs.SwaggerInfo.Title = "Shipping Address Microservice"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	endPoint := fmt.Sprintf(":%s", config.Server["PORT"])

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	go startGrpcServer()

	log.Printf("start http server listening %s", endPoint)

	_ = server.ListenAndServe()
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, &grpc_services.Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
