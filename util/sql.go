package util

import (
	"database/sql"
	"dbo-test/config"
	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Configure.Postgres.Address)
	if err != nil {
		return nil, err
	}
	return db, nil
}
