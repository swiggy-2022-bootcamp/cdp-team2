package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
)

// @title Checkout Api
// @version 1.0
// @description This is checkout service.

// @host 35.84.28.237:30206
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Start() {
	router, err := InitRouter()

	if err != nil {
		log.Println("Error starting server")
		return
	}

	endPoint := fmt.Sprintf(":%s", config.Config["PORT"])

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	log.Printf("start http server listening %s", endPoint)

	_ = server.ListenAndServe()
}
