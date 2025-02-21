package main

import "database/sql"

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
