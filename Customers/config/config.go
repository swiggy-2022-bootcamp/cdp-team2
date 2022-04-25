package config

var Server = map[string]string{
	"PORT": "8093",
	"AUTH_PORT":"",
	"GRPC_PORT":"8094",
	"ADDRESS_PORT":"9001",
}

var AWS = map[string]string{
	"TABLE_NAME":"team-2-Customers",
	"region":"us-west-2",
	"endpoint":"https://dynamodb.us-west-2.amazonaws.com",
}


