##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o categories-app cmd/server/main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/categories-app /categories-app

EXPOSE 7450

ENTRYPOINT ["/categories-app"]
