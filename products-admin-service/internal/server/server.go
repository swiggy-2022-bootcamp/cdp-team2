package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/products-admin-service/config"
	"github.com/products-admin-service/internal/core/ports"
)

type Server struct {
	Handlers ports.IHandlers
}

func NewServer(handlers ports.IHandlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

func (s *Server) Initialize() {
	server := gin.Default()

	productsRoutes := server.Group("/products/")
	productsRoutes.GET("/", s.Handlers.Health)

	log.Fatal(server.Run(config.Config["PORT"]))
}
