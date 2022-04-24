package config

var Server = map[string]string{
	"PORT": "8003",
}

var AWS = map[string]string{
	"region":   "us-west-2",
	"endpoint": "http://localhost:8000",
}

var GrpcAdd = map[string]string{
	"ORDER_SERVICE": "0.tcp.in.ngrok.io:19096",
}
