##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o checkout-app cmd/server/main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/checkout-app /checkout-app

EXPOSE 7451

ENTRYPOINT ["/checkout-app"]
