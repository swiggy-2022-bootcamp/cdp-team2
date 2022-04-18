package httphandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

// AddReward godoc
// @Summary Add reward points to the customer
// @Description This API allows Admin to directly add reward points to the customer. It accepts reward points and description in the request body, and on successful insertion in database, it returns 200 OK.
// @Tags healthcheck
// @Accept json
// @Success 200 {object} models.Reward
// @Failure 500
// @Router /reward/v1/reward/{id} [post]
func AddRewardHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//customerId := util.ExtractDetailsFromToken(req)

		// read request body
		b, goErr := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while reading the request body")
			http.Error(w, errors.InternalError.ErrorMessage, errors.InternalError.HttpResponseCode)
			return
		}

		// unmarshal the request to Product model
		var reward models.Reward

		goErr = json.Unmarshal(b, &reward)
		if goErr != nil {
			log.WithError(goErr).Error("an error occurred while unmarshalling the request")
			http.Error(w, errors.UnmarshalError.ErrorMessage, errors.UnmarshalError.HttpResponseCode)
			return
		}

		// reward.CustomerId = customerId
		service := services.GetAddRewardService()

		// Process the request
		err := service.ProcessRequest(reward)
		if err != nil {
			http.Error(w, err.ErrorMessage, err.HttpResponseCode)
			return
		}

		w.WriteHeader(200)
	}
}
