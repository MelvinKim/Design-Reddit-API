package config

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	DSN string
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() (*sql.DB, error) {
	flag.StringVar(&DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=reddit-api sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	connection, err := openDB(DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")

	return connection, nil
}
