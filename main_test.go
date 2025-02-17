package main_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Postgres()(context.Context, *postgres.PostgresContainer, error){
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

func TestMain(t *testing.T) {
	log.Print("not implemented")
	ctx, pc, err := Postgres()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		pc.Terminate(ctx)
	}()

}
