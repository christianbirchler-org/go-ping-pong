package main

import (
	"log/slog"
	"net/http"
)

type EchoHandler struct {
}

func (h *EchoHandler)ServeHTTP(rw http.ResponseWriter, r *http.Request){
	slog.Info("echo handler", "request", r, "response", rw)
	rw.Write([]byte("pong"))
}


func main() {
	slog.Info("start go-sample")

	http.Handle("/ping", &EchoHandler{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("server failed", "error", err)
	}
}
