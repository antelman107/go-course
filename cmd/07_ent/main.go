package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/antelman107/go-course/internal/ent"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	connectionString := "host=localhost user=course password=course dbname=course sslmode=disable"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()

	// Run the auto migration tool.
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

	e, err := client.User.Get(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.ID, e.Name)

	all, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range all {
		fmt.Println(item.String())
	}
}
