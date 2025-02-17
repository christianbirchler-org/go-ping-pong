package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

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

	connStr := "postgresql://..."
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("connect to DB", err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("get driver with db instance", err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	if err != nil {
		log.Fatal("create migration of the DB", err.Error())
	}
	err = m.Up()
	if err != nil {
		log.Fatal("migrate DB", err.Error())
	}

	http.Handle("/ping", &PingHandler{
		counter: &PostgresCounter{
			db: db,
		},
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("server failed", "error", err)
	}
}
