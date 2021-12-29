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

	var (
		id   int
		name string
	)

	err = db.QueryRow("SELECT id, name FROM users WHERE id = $1", 1).Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)
}
