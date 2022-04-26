package config

var Server = map[string]string{
	"PORT": "7450",
}

var AWS = map[string]string{
	"region":   "us-west-2",
	"endpoint": "http://localhost:8000",
}

var base = "35.84.28.237"

var GrpcAdd = map[string]string{
	"CART_SERVICE":    base + ":30217",
	"ORDER_SERVICE":   base + ":30208",
	"PRODUCT_SERVICE": base + ":30202",
	"REWARD_SERVICE":  base + ":30219",
}
