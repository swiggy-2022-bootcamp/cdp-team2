package server

import (
	"log"
	"net"

	pb "common/protos/products"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/config"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"google.golang.org/grpc"
)

type Server struct {
	Handlers ports.IHandlers
}

func NewServer(handlers ports.IHandlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

// @title Swagger Products Admin Microservice
// @version 1.0
// @description Micorservice for handling admin products.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /products
func (s *Server) Initialize() {
	server := gin.Default()

	productsRoutes := server.Group("/products/")
	productsRoutes.GET("/", s.Handlers.Health)
	productsRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go StartGrpcServer()

	log.Fatal(server.Run(config.Config["PORT"]))
}

type GrpcServer struct {
	pb.UnimplementedProductsServicesServer
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":7500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductsServicesServer(s, &GrpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
