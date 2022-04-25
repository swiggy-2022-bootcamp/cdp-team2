package config

var Server = map[string]string{
	"PORT": "7450",
}

var AWS = map[string]string{
	"region":   "us-west-2",
	"endpoint": "http://localhost:8000",
}

var GrpcAdd = map[string]string{
	"CART_SERVICE":    "10.50.2.224:30202",
	"ORDER_SERVICE":   "10.50.2.224:30202",
	"PRODUCT_SERVICE": "10.50.2.224:30202",
	"REWARD_SERVICE":  "10.50.2.224:30202",
}
