package httphandlers

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

// HealthCheck godoc
// @Summary Checks the health of the microservice
// @Description Returns json stating with pass status and http response code 200 OK, otherwise returns http response code 500 Internal Server Error.
// @Tags healthcheck
// @Produce  json
// @Success 200
// @Failure 500
// @Router /reward/v1/healthcheck [get]
func HealthCheckHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetHealthCheckService()

		// Process the request
		err := service.ProcessRequest()
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
