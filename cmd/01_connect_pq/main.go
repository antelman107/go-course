package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := "user=course password=course dbname=course sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
