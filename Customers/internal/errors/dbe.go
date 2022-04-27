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

	CartEmptyError = ServerError{
		ErrorMessage:     "no products found in cart",
		HttpResponseCode: http.StatusInternalServerError,
	}

	ProductNotFoundInCart = ServerError{
		ErrorMessage:     "product not found in cart",
		HttpResponseCode: http.StatusNotFound,
	}

	ProductNotExistError = ServerError{
		ErrorMessage:     "product with given product id does not exist",
		HttpResponseCode: http.StatusNotFound,
	}

	EmptyCartFailedError = ServerError{
		ErrorMessage:     "an error occurred while deleting the cart",
		HttpResponseCode: http.StatusInternalServerError,
	}
)
