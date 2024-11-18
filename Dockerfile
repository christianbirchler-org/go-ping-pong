FROM golang:1.23 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ping-pong main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/ping-pong /app

CMD ["/app/ping-pong"]
