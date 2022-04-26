package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

// HealthCheck godoc
// @Summary Checks the health of the microservice
// @Description Returns json stating with pass status and http response code 200 OK, otherwise returns http response code 500 Internal Server Error.
// @Tags Cart
// @Produce  json
// @Success 200 {object} models.HealthCheck
// @Failure 500
// @Router /cart/v1/healthcheck [get]
func HealthCheckHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetHealthCheckService()

		// Process the request
		response := service.ProcessRequest()

		respBytes, err := json.Marshal(response)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError)
			return
		}

		_, err = w.Write(respBytes)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError)
			return
		}
	}
}

func writeResponse(w http.ResponseWriter, status int) {
	downResponse := models.HealthCheck{Status: "Down"}
	respBytes, _ := json.Marshal(downResponse)

	w.WriteHeader(status)
	w.Write(respBytes)
}
