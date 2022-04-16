package server

import (
	"fmt"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/config"
	"log"
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

	_ = server.ListenAndServe()
}
