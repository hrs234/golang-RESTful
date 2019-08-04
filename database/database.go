package database

import (
	"database/sql"
	"log"
)

// connection of databases
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/DBlogistik")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}

	return db
}
