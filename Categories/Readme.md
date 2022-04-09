Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:7450/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/categories .
```

docker run
```sh
docker run -p 7450:7450 --name categories cdp-team2/categories 
```

docker sh into container
```sh
docker exec -it categories /bin/sh
```