package main

import (
	"log/slog"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
)

type Counter interface {
	increment() (int, error)
	reset() (int, error)
}

type PostgresCounter struct {
	db *sql.DB
}

func (pc *PostgresCounter) increment() (int, error) {
	// TODO
	return 0, nil
}

func (pc *PostgresCounter) reset() (int, error) {
	// TODO
	return 0, nil
}

type PingHandler struct {
	counter Counter
}

func (h *PingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	slog.Debug("echo handler", "request", r, "response", rw)
	h.counter.increment()
	rw.Write([]byte("pong"))
}

func main() {
	slog.Info("start go-sample")

	connStr := ""
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("connect to DB", err)
	}

	_ = &PostgresCounter{
		db: db,
	}

	http.Handle("/ping", &PingHandler{
		counter: &PostgresCounter{},
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("server failed", "error", err)
	}
}
