package database

import (
	"os"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func Init() pg.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		DB = pg.Connect(&pg.Options{User: "postgres"})
	} else {
		DB = pg.Connect(&pg.Options{
			User:     "postgres",
			Database: "postgres",
			Addr:     dbURL,
		})
	}
	return *DB
}
