# Products-Admin Microservice
Products-Admin microservice is mainly responsible for
adding new products by an admin, update a product by its
ID, delete a product using ID, and search products based
upon it's categoryID, manufacturerID, keywords of product's
model name, sku(stock keeping unit), tags.
This Microservice also works as gRPC server, which runs on port
8002 for inter-microservice communication.

## Checklist
- [x] Completed Writing Microservice Assigned
- [x] Written Test Cases
- [x] Dockerizing an application
- [x] Swagger Documentation
- [x] Effective Use of GRPC
- [x] Container Orchestration
- [x] Have maintained in code repository
- [x] Builds the application using CI/CD
- [x] Data base model properly show cased


## Steps to setup locally

### Using Published DockerHub Image
```
# Pull image from dockerhub
> docker pull swiggyteam2/products-admin

# Run container from ahove image
> docker run -p 8001:8001 -p 8002:8002
```
### Using Docker
```
# Build docker Image
> docker build -t cdp-team2/products-admin-service .

# Run docker image
> docker run -p 8001:8001 -p 8002:8002 -d --name products-admin-service cdp-team2/products-admin-service

# Sh into Docker Container
> docker exec -it products-admin-service /bin/sh 
```

### Without Docker
```
# Download golang dependencies
> go mod download

# Start the CMD server
> go run cmd/server/main.go
```

### Swagger Docs
```
# Open the below url
http://localhost:8000/products/swagger/index.html

# Command to regenerate swagger docs
> swag init -g ./internal/server/server.go
```

### Endpoints Covered
| EndPoints Name | Endpoints | Tag | Method |
| -------------- | --------- | --- | -------|
| Add Product    | /api/rest_admin/products/ | Products | POST |
| Get All Products    | /api/rest_admin/products/ | Products | GET  |
| Update Product    | /api/rest_admin/products/{id} | Products | PATCH |
| Delete Product    | /api/rest_admin/products/{id} | Products | DELETE |
| Search Product By CategoryID    | /api/rest_admin/products/search/category/{id} | Search | GET |
| Search Product By ManufacturerID    | /api/rest_admin/products/search/manufacturer/{id} | Search | GET |
| Search Product By keyword    | /api/rest_admin/products/search/keyword/{keyword} | Search | GET |

### Written Test Cases
#### Test Cases written for Handlers methods

coverage: 87.5% of statements
![image](https://user-images.githubusercontent.com/64790109/165352747-bf198621-c920-4cd3-9fa2-13b93058f94b.png)

#### Test Cases written for Services methods
coverage: 91.3% of statements

![image](https://user-images.githubusercontent.com/64790109/165353284-f7357858-dbc6-49a9-8682-ed7a79045ca1.png)

#### DataModel Used
![image](https://user-images.githubusercontent.com/64790109/165353852-bfcc2ec5-5abc-4d37-844d-489d481dcd25.png)

#### Grpc Methods for inter-microservice communication

- GetAvailableProducts --> Called by Products-FrontStore MS
-	CheckProductsWithCategory --> Called by Category MS
-	IsProductExists --> Called by Category MS/ Checkout MS
-	GetProductById --> Called by Products-FrontStore/ Checkout MS
-	GetProductsByCategoryId --> Called by Products-FrontStore
-	DeductProductsQuantity --> Called by Checkout MS

#### Kubernetes minikube running services for products-admin

![image](https://user-images.githubusercontent.com/64790109/165355509-c4f1f93d-5f7b-445b-94ac-e51cdda8b459.png)

