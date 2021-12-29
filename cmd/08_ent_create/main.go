package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/antelman107/go-course/internal/ent"
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "host=localhost user=course password=course dbname=course sslmode=disable"
	client, err := ent.Open(dialect.Postgres, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	u, err := client.User.
		Create().
		SetName("fourth").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user was created: ", u)
}
