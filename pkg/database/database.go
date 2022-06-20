package database

import (
	"log"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func Init() pg.DB {
	DB = pg.Connect(&pg.Options{User: "postgres"})
	if DB == nil {
		log.Panic("Database connection error")
	}
	log.Println("Database connected")
	return *DB
}
