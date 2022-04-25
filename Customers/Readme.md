Customers Admin

-------------------------------------------------------
Architecture Diagram
====================
</br>

 <img src="https://github.com/swiggy-2022-bootcamp/cdp-team2/blob/customers-ms/Customers/images/arc1.PNG" width="800" height="500">
</br>
</br>

 <img src="https://github.com/swiggy-2022-bootcamp/cdp-team2/blob/customers-ms/Customers/images/arc2.PNG" width="800" height="500">
</br>
</br>

 <img src="https://github.com/swiggy-2022-bootcamp/cdp-team2/blob/customers-ms/Customers/images/arc3.PNG" width="800" height="500">
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
Features included in this Project:</br>
=================================</br>
1) CRUD using dynaomoDB
2) gRPC for service to service communication
3) Get Address details of the Customer
4) Send Address details of the Customer to Address service using gRPC
-------------------------------------------------------


API Requests:</br>
=============</br>

1) POST    /customers                          => Creates a new Customer  </br>
2) GET    /customers/{id}                       => Gets the detials of the customer along with address </br>
3) PUT    /customers/{id}                       => Updates a customer </br>
4) GET    /customers/email/{email}              => Gets the customer details using email</br>
5) DELETE /customers/{id}                       => Deletes customers using id
==================================</br>
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

