package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connection := "user=postgres dbname=dream_store password=250816 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
