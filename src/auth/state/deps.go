package state

import (
	"boilerplate/database"
	"boilerplate/guard"
	"boilerplate/users"
)

type Deps interface {
	DB() *database.DB
	GuardClient() guard.Client
	UsersClient() users.Client
}

type Dependencies struct {
	db          *database.DB
	guardClient guard.Client
	usersClient users.Client
}

var _ Deps = (*Dependencies)(nil)

func MakeDependencies(db *database.DB,
	usersClient users.Client, guardClient guard.Client) *Dependencies {
	return &Dependencies{
		db:          db,
		guardClient: guardClient,
		usersClient: usersClient,
	}
}

func (d *Dependencies) DB() *database.DB {
	return d.db
}

func (d *Dependencies) GuardClient() guard.Client {
	return d.guardClient
}

func (d *Dependencies) UsersClient() users.Client {
	return d.usersClient
}
