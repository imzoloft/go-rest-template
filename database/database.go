package database

import (
	"database/sql"
	"log"

	"github.com/imzoloft/go-rest-api/config"
	_ "github.com/lib/pq"
)

func NewSQLDatabase(cfg config.Database) *sql.DB {
	// change the database driver or use orm here
	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func HealthCheck(db *sql.DB) error {
	return db.Ping()
}
