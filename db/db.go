package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBConnect() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "host=localhost port=1111 user=postgresuser password=postgrespassword dbname=postgres sslmode=disable")
	if err != nil {
		println(`in loger`)
		log.Fatalln(err)
	}
	return db
}
