package db

import (
	"github.com/jmoiron/sqlx"
)

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		panic(err)
	} //Connecting to database
	return db
}
