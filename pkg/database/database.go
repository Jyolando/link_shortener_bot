package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "shortener_link_db"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetLink(db *sql.DB, link string) string {
	var id int
	var shortLink string
	var fullLink string

	rows, err := db.Query("SELECT * FROM links WHERE shortLink = $1", link)
	CheckError(err)
	for rows.Next() {
		err = rows.Scan(&id, &fullLink, &shortLink)
		CheckError(err)
	}
	if fullLink == "" {
		return "nil"
	} else {
		return fullLink
	}
}

func Init() *sql.DB {
	pgsqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", pgsqlconn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	log.Println("Database connection established")
	return db
}
