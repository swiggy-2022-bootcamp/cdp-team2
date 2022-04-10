Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:8080/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/customers .
```

docker run
```sh
docker run -p 8080:8080 --name customers cdp-team2/customers
```

docker sh into container
```sh
docker exec -it customers /bin/sh