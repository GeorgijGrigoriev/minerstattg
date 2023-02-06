package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Get() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=tgbot sslmode=disabled")
	if err != nil {
		log.Fatal(err.Error())
	}

	db.MustExec(schema)
}
