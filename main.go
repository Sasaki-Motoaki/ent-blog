package main

import (
	"context"
	"example/ent-blog/ent"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=db user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// if user, err := CreateUser(context.Background(), client); err != nil {
	// 	fmt.Println(user)
	// }
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetEmail("example@example.com").
		SetPassword("password").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
