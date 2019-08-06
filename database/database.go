package database

import (
	"database/sql"
	"log"

	"os"

	"github.com/joho/godotenv"
)

// connection of databases
func Connect() *sql.DB {
	checkENV := godotenv.Load()
	if checkENV != nil {
		log.Fatal("Failed Load ENV file!")
		log.Fatal(checkENV)
	}

	dbName := os.Getenv("DATABASE")
	dbPort := os.Getenv("DATABASE_PORT")
	host := os.Getenv("HOST")

	db, err := sql.Open("mysql", "root:@tcp("+host+":"+dbPort+")/"+dbName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}

	return db
}
