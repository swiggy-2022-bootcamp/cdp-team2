Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:8081/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/customer-account .
```

docker run
```sh
docker run -p 8081:8081 --name customer-account cdp-team2/customer-account
```

docker sh into container
```sh
docker exec -it customer-account /bin/sh