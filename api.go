package main

import (
	"log/slog"
	"net/http"
)

type PingHandler struct {
	counter Counter
}

func (h *PingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	slog.Debug("echo handler", "request", r, "response", rw)
	h.counter.increment()
	rw.Write([]byte("pong"))
}
