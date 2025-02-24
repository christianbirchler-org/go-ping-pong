package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

func NewPostgresContainer() (context.Context, *postgres.PostgresContainer, error) {
	ctx := context.Background()

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	return ctx, postgresContainer, err
}

func NewPostgresDatabase(ctx context.Context, t *testing.T, pc *postgres.PostgresContainer) (*sql.DB, error) {
	h, err := pc.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	p, err := pc.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"
	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + h + ":" + p.Port() + "/" + dbName + "?sslmode=disable"

	return sql.Open("postgres", connStr)

}

func NewDatabaseMigration(db *sql.DB, t *testing.T) (*migrate.Migrate, error) {
	driver, err := pg.WithInstance(db, &pg.Config{})
	if err != nil {
		t.Fatal("get driver with db instance", err.Error())
	}

	return migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
}

func TestE2ESinglePing(t *testing.T) {
	ctx, pc, err := NewPostgresContainer()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		pc.Terminate(ctx)
	}()
	err = pc.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = pc.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}

	db, err := NewPostgresDatabase(ctx, t, pc)
	if err != nil {
		t.Fatal(err)
	}

	m, err := NewDatabaseMigration(db, t)
	err = m.Up()
	if err != nil {
		t.Fatal("migrate DB", err.Error())
	}

	// test subject
	c := &PostgresCounter{db: db}
	ph := &PingHandler{counter: c}

	ts := httptest.NewServer(ph)
	defer ts.Close()

	// task
	tc := ts.Client()
	resp, err := tc.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}
	var pongResponse PongResponse
	err = json.NewDecoder(resp.Body).Decode(&pongResponse)
	if err != nil {
		t.Fatal(err)
	}

	// oracle
	if pongResponse.Msg != "pong" {
		t.Fatal("pong response:", pongResponse)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("response status code:", resp.StatusCode)
	}

}
