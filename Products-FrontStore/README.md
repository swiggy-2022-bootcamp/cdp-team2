## Products-FrontStore Microservice

### Steps to setup locally

#### Using Docker
```
# Build docker Image
> docker build -t cdp-team2/products-frontstore-service .

# Run docker image
> docker run -p 8001:8001 -d --name products-frontstore-service cdp-team2/products-frontstore-service

# Sh into Docker Container
> docker exec -it products-frontstore-service /bin/sh 
```

#### Without Docker
```
# Download golang dependencies
> go mod download

# Start the CMD server
> go run cmd/server/main.go
```

#### Swagger Docs
```
# Open the below url
http://localhost:8001/api/rest/products/swagger/index.html

# Command to regenerate swagger docs
> swag init ./internal/server/server.go
```
