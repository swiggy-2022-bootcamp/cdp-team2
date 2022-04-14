package httphandlers

import (
	"encoding/json"
	"net/http"

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
