# Order Microservice - Taranjeet Singh

## Architecture

![image](https://user-images.githubusercontent.com/42766421/164993173-588a7905-0898-491e-b0e5-7c08ca52beb5.png)

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
1) CRUD for orders using Admin
2) gRPC for service to service communication
3) Create order and update order status from the checkout service gRPC calls
4) Update the total payed price from the reward microservice gRPC calls
5) Update the delivery address from the shipping microservice gRPC calls
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
<br>- ![image](https://user-images.githubusercontent.com/42766421/164993557-5e988dea-8957-4ed1-b2ab-e5b8774c2a75.png)


Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:8002/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/order-ms .
```

docker run
```sh
docker run -p 8080:8080 --name order-ms cdp-team2/order-ms
```

docker sh into container
```sh
docker exec -it order-ms /bin/sh

