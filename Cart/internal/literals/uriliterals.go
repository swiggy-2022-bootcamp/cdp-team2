package literals

const (
	APIVersion1 = "/v1"

	HealthCheckAPIName  = "Health Check API"
	HealthCheckEndpoint = APIVersion1 + "/healthcheck"

	GetCartAPIName  = "Get Cart API"
	GetCartEndpoint = APIVersion1 + "/cart/{customerId}"

	AddCartAPIName      = "Add Cart API"
	AddCartItemEndpoint = APIVersion1 + "/cart"

	UpdateCartItemAPIName  = "Update Cart Item API"
	UpdateCartItemEndpoint = APIVersion1 + "/cart"

	DeleteCartItemAPIName  = "Delete Cart Item API"
	DeleteCartItemEndpoint = APIVersion1 + "{customerId}/cart/{productId}"
)
