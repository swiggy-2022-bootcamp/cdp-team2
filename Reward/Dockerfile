FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o reward-app cmd/server/main.go

FROM alpine

WORKDIR /

COPY --from=build /app/reward-app /reward-app

EXPOSE 8001

ENTRYPOINT ["/reward-app"]
