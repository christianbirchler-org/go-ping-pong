FROM golang:1.23 AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-ping-pong .

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/go-ping-pong /app
COPY --from=build /app/migrations /app/migrations

CMD ["/app/go-ping-pong"]
