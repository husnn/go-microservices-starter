//go:build logical

package state

import (
	"context"
	"github.com/rs/zerolog/log"
	authLogical "boilerplate/auth/client/logical"
	"boilerplate/database"
	guardLogical "boilerplate/guard/client/logical"
	usersLogical "boilerplate/users/client/logical"
	usersDB "boilerplate/users/db"
)

func New() deps {
	db, err := database.Connect(
		context.Background(), usersDB.SeedInternal)
	if err != nil {
		log.Fatal().Err(err).
			Msg("error connecting to database")
	}

	var d dependencies

	d.db = db
	d.guardClient = guardLogical.NewClient(&d)
	d.usersClient = usersLogical.NewClient(&d)
	d.authClient = authLogical.NewClient(&d)

	return &d
}

type deps interface {
	Deps
}

type dependencies struct {
	Dependencies
}

var _ deps = (*dependencies)(nil)
