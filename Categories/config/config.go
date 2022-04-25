package config

var Server = map[string]string{
	"PORT": "7450",
}

var AWS = map[string]string{
	"region":   "us-west-2",
	"endpoint": "http://localhost:8000",
}

var GrpcAdd = map[string]string{
	"CART_SERVICE":    "localhost:7550",
	"ORDER_SERVICE":   "localhost:7551",
	"PRODUCT_SERVICE": "localhost:7552",
	"REWARD_SERVICE":  "localhost:7553",
}
