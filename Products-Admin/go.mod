module github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.1
)

require (
	common v0.0.0-00010101000000-000000000000
	github.com/swaggo/swag v1.7.9
	google.golang.org/grpc v1.45.0
)

replace common => ../common
