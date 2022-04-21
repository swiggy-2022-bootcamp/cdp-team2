FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o account-app cmd/server/main.go

FROM alpine

WORKDIR /

COPY --from=build /app/account-app /account-app

EXPOSE 8050

ENTRYPOINT ["/account-app"]
