package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type PingHandler struct {
	counter Counter
}

type PongResponse struct {
	Msg string `json:"msg"`
}

func (h *PingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	slog.Debug("echo handler", "request", r, "response", rw)
	h.counter.increment()

	pr := PongResponse{Msg: "pong"}

	json.NewEncoder(rw).Encode(pr)
}
