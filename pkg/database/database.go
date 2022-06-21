package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	UserId  string
	Attempt int
}

var (
	db  *sqlx.DB
	err error
)

func Init() {
	db, err = sqlx.Connect("postgres", "user=postgres dbname=* password=* sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	test := []User{}
	err = db.Select(&test, `SELECT userid, attempt FROM users`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(test)
}
