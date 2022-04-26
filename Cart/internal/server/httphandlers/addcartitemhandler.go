package httphandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"

	log "github.com/sirupsen/logrus"
)

// AddCartItem godoc
// @Summary Add the cart item in the cart.
// @Description It accepts product id and quantity in the request url. If the product id exists, then it adds to the cart otherwise return 404. On successful addition of item to the cart, returns 200 OK, otherwise 500 Internal Server Error.
// @Tags Cart
// @Produce  json
// @Success 200
// @Failure 500
// @Router /cart/v1/cart [post]
func AddCartItemHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		customerId := util.ExtractDetailsFromToken(req)
		// read request body
		b, goErr := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while reading the request body")
			http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
			return
		}

		// unmarshal the request to Product model
		var product models.Product
		goErr = json.Unmarshal(b, &product)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while unmarshalling the request")
			http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		service := services.GetAddCartItemService()

		// Process the request
		err := service.ProcessRequest(customerId, product)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
