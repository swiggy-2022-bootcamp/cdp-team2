# Builder Image
FROM golang:1.17 as builder

# Create directory for common
WORKDIR /common

# Copy common 
COPY /common .

# Create directory for main app
WORKDIR /app

# Install dependecies
COPY /Products-Admin/go.mod .
COPY /Products-Admin/go.sum .
RUN go mod download


# Copy all files to app
COPY /Products-Admin .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./build/main ./cmd/server/main.go

# LightWeight Final Image
FROM alpine:latest
COPY --from=builder app/build/main .

EXPOSE 8001
EXPOSE 8002
# Executable
ENTRYPOINT ["./main"]
