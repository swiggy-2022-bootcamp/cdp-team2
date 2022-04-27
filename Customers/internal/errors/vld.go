package errors

import (
	"net/http"

	"customers/internal/literals"
)

// Validation errors
var (
	ParametersMissingError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.ParametersMissing,
		HttpResponseCode: http.StatusBadRequest,
	}

	QueryParamMissingError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.QueryParamMissing,
		HttpResponseCode: http.StatusBadRequest,
	}

	MalformedQueryParamError = ServerError{
		ErrorMessage:     literals.AppPrefix + literals.MalformedQueryParam,
		HttpResponseCode: http.StatusBadRequest,
	}
)
