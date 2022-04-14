package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

// UpdateCartItem godoc
// @Summary Updates the cart item count
// @Description It accepts product id and updated count in request body, and return 200 OK in case of successful update. Otherwise return 500 Internal Server Error, or 404 if the product id is not present in cart.
// @Produce  json
// @Success 200
// @Failure 500
// @Router /cart [put]
func UpdateCartItemHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetUpdateCartItemService()

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
