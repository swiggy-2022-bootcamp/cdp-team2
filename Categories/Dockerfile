##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /common
COPY common/ .

WORKDIR /app

COPY Categories/go.mod .
COPY Categories/go.sum .
RUN go mod download

COPY Categories/ .

RUN go build -o categories-app cmd/server/main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/categories-app /categories-app

EXPOSE 7450

ENTRYPOINT ["/categories-app"]
