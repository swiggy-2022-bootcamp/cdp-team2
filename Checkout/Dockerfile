##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /common
COPY common/ .

WORKDIR /app
COPY Checkout/go.mod .
COPY Checkout/go.sum .
RUN go mod download

COPY Checkout/ .

RUN go build -o checkout-app cmd/server/main.go

# ENTRYPOINT ["/bin/sh"]

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/checkout-app /checkout-app

EXPOSE 7451

ENTRYPOINT ["/checkout-app"]
