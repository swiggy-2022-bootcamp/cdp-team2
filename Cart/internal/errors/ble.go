package errors

import "net/http"

// ServerError defines the error message and related http response code for any
// error return by the microservice.
type ServerError struct {
	ErrorMessage     string
	HttpResponseCode int
}

var (
	InternalError = ServerError{
		ErrorMessage:     "an error occurred while handling the request",
		HttpResponseCode: http.StatusInternalServerError,
	}

	MarshalError = ServerError{
		ErrorMessage:     "an error occured while marshalling the response",
		HttpResponseCode: http.StatusInternalServerError,
	}

	UnmarshalError = ServerError{
		ErrorMessage:     "an error occured while unmarshalling the request",
		HttpResponseCode: http.StatusInternalServerError,
	}
)
