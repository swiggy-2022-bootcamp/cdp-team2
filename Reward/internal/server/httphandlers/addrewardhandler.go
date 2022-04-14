package httphandlers

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

// AddReward godoc
// @Summary Add reward points to the customer
// @Description This API allows Admin to directly add reward points to the customer. It accepts reward points and description in the request body, and on successful insertion in database, it returns 200 OK.
// @Tags healthcheck
// @Produce  json
// @Success 200
// @Failure 500
// @Router /reward/{id} [post]
func AddRewardHandler(config *util.RouterConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		service := services.GetAddRewardService()

		// Process the request
		service.ProcessRequest()

		w.WriteHeader(200)
	}
}
