package config

var Config = map[string]string{
	"PORT": "7451",
}

var base = "35.84.28.237"

var GrpcAdd = map[string]string{
	"CART_SERVICE":    base + ":30217",
	"ORDER_SERVICE":   base + ":30208",
	"PRODUCT_SERVICE": base + ":30201",
	"REWARD_SERVICE":  base + ":30219",
}
