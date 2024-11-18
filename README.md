# go-ping-pong
This is a sample project for bootstrapping a web application with Go.

## Quick Start

``` shell
go run main.go
```

``` shell
curl http://localhost:8080/ping
```

You should get a HTTP response including a `pong`.

## Usage with Docker

``` shell
docker run --rm -p 8080:8080 ghcr.io/christianbirchler-org/go-ping-pong:latest
```

``` shell
curl http://localhost:8080/ping
```

## License
[MIT](./LICENSE.txt)
