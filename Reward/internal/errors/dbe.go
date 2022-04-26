package errors

import "net/http"

var (
	RecordNotFound = ServerError{
		ErrorMessage:     "record not found",
		HttpResponseCode: http.StatusNotFound,
	}

	DatabaseQueryFailed = ServerError{
		ErrorMessage:     "database query failed",
		HttpResponseCode: http.StatusInternalServerError,
	}

	MarshalToAttributesError = ServerError{
		ErrorMessage:     "marshalling Go value type to a map of AttributeValues failed",
		HttpResponseCode: http.StatusInternalServerError,
	}

	UnmarshalAttributesError = ServerError{
		ErrorMessage:     "unmarshalling from a map of AttributeValues failed",
		HttpResponseCode: http.StatusInternalServerError,
	}
)
