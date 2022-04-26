module github.com/swiggy-2022-bootcamp/cdp-team2/Shipping

go 1.16

require (
	github.com/aws/aws-sdk-go v1.43.41
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/golang/mock v1.6.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/gin-swagger v1.4.2
	github.com/swaggo/swag v1.8.1
	github.com/ugorji/go v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220421235706-1d1ef9303861 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
)

replace common => ../common
