package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	connectionString := "host=localhost user=course password=course dbname=course sslmode=disable"
	ctx := context.Background()
	db, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}
}
