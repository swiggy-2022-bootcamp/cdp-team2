# Builder Image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./build/main ./cmd

# Final Image
FROM alpine:latest
COPY --from=builder app/build/main .

# Executable
ENTRYPOINT ["./main"]
