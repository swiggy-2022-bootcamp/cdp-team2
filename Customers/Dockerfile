FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o customer-app cmd/server/main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/customer-app /customer-app

EXPOSE 8050
EXPOSE 9000

ENTRYPOINT ["/customer-app"]
