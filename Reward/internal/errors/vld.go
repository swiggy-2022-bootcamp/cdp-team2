package errors

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/literals"
)

var (
	ParametersMissingError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.ParametersMissing,
		HttpResponseCode: http.StatusBadRequest,
	}
)
