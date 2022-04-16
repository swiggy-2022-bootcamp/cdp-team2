package httphandlers

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

// DeleteCartItem godoc
// @Summary Deletes the cart item
// @Description Deletes the product from the cart of user. Product id is given as path parameter in request URL. Status 200, returned on successful deletion with no content.
// @Tags healthcheck
// @Produce  json
// @Success 200
// @Failure 500
// @Router /cart/{key} [delete]
func DeleteCartItemHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetDeleteCartItemService()

		// Process the request
		err := service.ProcessRequest()
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
