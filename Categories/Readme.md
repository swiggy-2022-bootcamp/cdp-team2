Run locally 
```sh
AWS_ACCESS_KEY_ID=fake AWS_SECRET_ACCESS_KEY=fake go run cmd/server/main.go
```

Generate Mocks
```sh
go generate ./...
```

Run Tests
```sh
go test  -v -coverprofile cover.out ./...

go tool cover -html=cover.out -o cover.html
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:7450/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/categories -f Dockerfile ../
```

docker run
```sh
docker run -p 7450:7450 --name categories cdp-team2/categories 
```

docker sh into container
```sh
docker exec -it categories /bin/sh
```


GRPCurl 
```
grpcurl -plaintext -d '{"category_id":10}' 127.0.0.1:7459 category.CategoryService/GetCategory
```