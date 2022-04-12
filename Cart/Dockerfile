FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o cart-app cmd/server/main.go

FROM alpine

WORKDIR /

COPY --from=build /app/cart-app /cart-app

EXPOSE 8000

ENTRYPOINT ["/cart-app"]
