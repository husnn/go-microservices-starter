package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"boilerplate/database"
	users_db "boilerplate/users/db"
)

func main() {
	ctx := context.Background()

	seedMock := func(ctx context.Context, dbc *pgxpool.Pool) error {
		return nil
	}

	_, err := database.Connect(ctx, users_db.SeedInternal, seedMock)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
}
