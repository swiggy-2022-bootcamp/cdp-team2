package config

var Config = map[string]string{
	"PORT": "7451",
}

var base = "35.84.28.237"

var GrpcAdd = map[string]string{
	"AUTH_SERVICE":    base + ":30220",
	"CART_SERVICE":    base + ":30217",
	"ORDER_SERVICE":   base + ":30208", //"0.tcp.in.ngrok.io:11488",
	"PRODUCT_SERVICE": base + ":30202",
	"REWARD_SERVICE":  base + ":30219",
}
