package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	slog.Info("start go-sample")

	_ = godotenv.Load(".env")
	connStr := os.Getenv("PG_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("connect to DB", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("get driver with db instance", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
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
