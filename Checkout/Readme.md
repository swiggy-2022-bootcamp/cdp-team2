Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:7451/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/checkout .
```

docker run
```sh
docker run -p 7451:7451 --name checkout cdp-team2/checkout 
```

docker sh into container
```sh
docker exec -it checkout /bin/sh
```