package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/config"
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

	endPoint := fmt.Sprintf(":%s", config.Config["PORT"])

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	log.Printf("start http server listening %s", endPoint)

	_ = server.ListenAndServe()
}
