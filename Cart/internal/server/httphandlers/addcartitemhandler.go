package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

// AddCartItem godoc
// @Summary Add the cart item in the cart.
// @Description It accepts product id and quantity in the request url. If the product id exists, then it adds to the cart otherwise return 404. On successful addition of item to the cart, returns 200 OK, otherwise 500 Internal Server Error.
// @Tags healthcheck
// @Produce  json
// @Success 200
// @Failure 500
// @Router /cart [post]
func AddCartItemHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetAddCartItemService()

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
