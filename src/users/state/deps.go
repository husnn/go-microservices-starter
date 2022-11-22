package state

import (
	"boilerplate/database"
	"boilerplate/guard"
)

type Deps interface {
	DB() *database.DB
	GuardClient() guard.Client
}

type Dependencies struct {
	db          *database.DB
	guardClient guard.Client
}

var _ Deps = (*Dependencies)(nil)

func MakeDependencies(db *database.DB,
	guardClient guard.Client) *Dependencies {
	return &Dependencies{
		db:          db,
		guardClient: guardClient,
	}
}

func (d *Dependencies) DB() *database.DB {
	return d.db
}

func (d *Dependencies) GuardClient() guard.Client {
	return d.guardClient
}
