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

	rows, err := db.Query("SELECT id, name FROM users WHERE id = $1", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, name)
	}

	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}
}
