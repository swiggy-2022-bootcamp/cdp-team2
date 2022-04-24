Customer-Account Frontstore

-------------------------------------------------------
Architecture Diagram
====================
</br>

 <img src="https://github.com/swiggy-2022-bootcamp/training-i-plus-plus/blob/main/uday/mini-project/MicroServices_ShoppingCartProject/aserviceProvider/diagram.PNG" width="800" height="500">
</br>


-------------------------------------------------------
Technologies/tools Used:
=================
    Golang
    gRPC
    SwaggerUI
    Microservices
    Dynamodb
    Gin
    SonarQube
    Docker
    Jenkins
    Kubernetes
    Mock testing

-------------------------------------------------------
Features included in this Project:
=================================
1) CRUD using dynaomoDB
2) gRPC for service to service communication
3) Get Cart details of the Customer
4) Get Reward details of the Customer
-------------------------------------------------------


API Requests:
=============

1) POST    /register                          => Creates a new Customer  </br>
2) GET    /account/{id}                       => Gets the detials of the customer along with Cart list and Reward </br>
3) PUT    /account/{id}                       => Updates a new User </br>

==================================
Starting Go Customer-Account-Fronststore server:</br>
-----------------------------
              go run cmd/server/main.go </br>


Run locally 
```sh
go run cmd/server/main.go
```

generate swagger docs
```sh
swag init -g ./internal/server/server.go
```

swagger page
> http://localhost:8050/swagger/index.html

docker build image
```sh
docker build -t cdp-team2/customer-account .
```

docker run
```sh
docker run -p 8050:8050 --name customer-account cdp-team2/customer-account
```

docker sh into container
```sh
docker exec -it customer-account /bin/sh



