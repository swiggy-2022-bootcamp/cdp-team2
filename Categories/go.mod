module github.com/swiggy-2022-bootcamp/cdp-team2/Categories

go 1.16

require (
	common v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.43.40
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.1
	github.com/swaggo/swag v1.7.9
	google.golang.org/grpc v1.45.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace common => ../common
