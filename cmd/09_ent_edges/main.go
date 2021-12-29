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

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	pet, err := client.Pet.
		Create().
		SetName("Some pet").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("pet was created: ", pet)

	group, err := client.Group.
		Create().
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("group was created: ", group)

	u, err := client.User.
		Create().
		SetName("fifth").
		AddPets(pet).
		AddPetIDs(pet.ID).
		AddGroups(group).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user was created: ", u)

	owners, err := client.Pet.QueryOwner(pet).All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range owners {
		fmt.Println(o)
	}

}


