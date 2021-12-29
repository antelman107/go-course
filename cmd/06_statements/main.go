package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=course password=course dbname=course sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO users (name) VALUES ($1)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("third")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rowsAffected)
}
