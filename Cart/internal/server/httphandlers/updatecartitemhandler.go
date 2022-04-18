package httphandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
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

		service := services.GetUpdateCartItemService()

		err := service.ValidateRequest(product)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		err = service.ProcessRequest("133", product)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
