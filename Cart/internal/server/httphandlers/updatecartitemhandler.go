package httphandlers

import (
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
		err := service.ProcessRequest()
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
