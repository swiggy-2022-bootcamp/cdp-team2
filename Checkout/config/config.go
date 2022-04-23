package config

var Config = map[string]string{
	"PORT": "7451",
}

var GrpcAdd = map[string]string{
	"CART_SERVICE":    "localhost:7550",
	"ORDER_SERVICE":   "localhost:7551",
	"PRODUCT_SERVICE": "localhost:7552",
	"REWARD_SERVICE":  "localhost:7553",
}
