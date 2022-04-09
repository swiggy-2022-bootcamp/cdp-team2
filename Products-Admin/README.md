## Products-Admin Microservie

### Steps to setup locally

#### Using Docker
```
# Build docker Image
> docker build -t cdp-team2/products-admin-service .

# Run docker image
> docker run -p 8000:8000 -d --name products-admin-service cdp-team2/products-admin-service

# Sh into Docker Container
> docker exec -it products-admin-service /bin/sh 
```

#### Without Docker
```
# Download golang dependencies
> go mod download

# Start the CMD server
> go run cmd/main.go
```

#### Swagger Docs
```
# Open the below url
http://localhost:8000/products/swagger/index.html

# Command to regenerate swagger docs
> swag init ./internal/server/server.go
```
