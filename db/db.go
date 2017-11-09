package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func DBConnect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=db user=pg-user password=password dbname=sample_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
