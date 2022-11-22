package state

import (
	"boilerplate/database"
	"boilerplate/users"
)

type Deps interface {
	DB() *database.DB
	UsersClient() users.Client
}

type Dependencies struct {
	db          *database.DB
	usersClient users.Client
}

var _ Deps = (*Dependencies)(nil)

func MakeDependencies(db *database.DB,
	usersClient users.Client) *Dependencies {
	return &Dependencies{
		db:          db,
		usersClient: usersClient,
	}
}

func (d *Dependencies) DB() *database.DB {
	return d.db
}

func (d *Dependencies) UsersClient() users.Client {
	return d.usersClient
}
