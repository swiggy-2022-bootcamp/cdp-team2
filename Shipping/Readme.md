# Shipping Address Microservice - Taranjeet Singh

## Architecture

![image](https://user-images.githubusercontent.com/42766421/164994074-732fd1a6-b756-4198-9737-7167fc7c38c0.png)

## Technologies/tools Used:
    Golang
    gRPC
    SwaggerUI
    Microservices
    Dynamodb
    Gin
    Docker
    Jenkins
    Kubernetes
    Mock testing

## Features included in this Project
1) CRUD for shipping address using Admin
2) gRPC for service to service communication
3) Set the address to an order with gRPC call to Order Microservice
-------------------------------------------------------


### Tasks

- [x] Completed Writing Microservice Assigned
- [x] Written Test Cases
- [x] Dockerizing an application
- [x] Swagger Documentation
- [x] Effective Use of GRPC
- [x] Container Orchestration
- [x] Have maintained in code repository
- [x] Data base model properly show cased


### Swagger
<br>- ![image](https://user-images.githubusercontent.com/42766421/164994126-c3dcfc80-abbf-47be-b6d1-e0acde8ea80b.png)



Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:8003/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/shipping-ms .
```

docker run
```sh
docker run -p 8080:8080 --name shippping-ms cdp-team2/shipping-ms
```

docker sh into container
```sh
docker exec -it shipping-ms /bin/sh

