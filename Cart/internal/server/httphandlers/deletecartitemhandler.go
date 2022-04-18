package httphandlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

// DeleteCartItem godoc
// @Summary Deletes the cart item
// @Description Deletes the product from the cart of user. Product id is given as path parameter in request URL. Status 200, returned on successful deletion with no content.
// @Tags Cart
// @Produce  json
// @Success 200
// @Failure 500
// @Router /cart/v1/cart/{key} [delete]
func DeleteCartItemHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		productId := params["key"]

		service := services.GetDeleteCartItemService()

		// Process the request
		err := service.ProcessRequest("133", productId)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)

		}

		w.WriteHeader(http.StatusOK)
	}
}
