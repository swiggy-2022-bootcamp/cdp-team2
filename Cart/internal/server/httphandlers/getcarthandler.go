package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

// GetCart godoc
// @Summary To get the cart details
// @Description It returns the cart details with status code 200 OK.
// @Produce  json
// @Success 200
// @Failure 500
// @Router /cart [get]
func GetCartHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetGetCartService()

		// Process the request
		response, err := service.ProcessRequest()
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		respBytes, goErr := json.Marshal(response)
		if goErr != nil {
			http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
			return
		}

		_, goErr = w.Write(respBytes)
		if goErr != nil {
			http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
			return
		}
	}
}
