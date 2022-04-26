package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/docs"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Order/protos/order"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/services/grpc_services"
	"google.golang.org/grpc"
)

// @title Order Microservice
// @version 1.0
// @description This is for creating orders during checkout.

// @host 35.84.28.237:30206
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Start() {
	router := InitRouter()

	docs.SwaggerInfo.Title = "Order Microservice"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
