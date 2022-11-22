package state

import (
	"boilerplate/auth"
	"boilerplate/database"
	"boilerplate/guard"
	"boilerplate/users"
)

type Deps interface {
	AuthClient() auth.Client
	DB() *database.DB
	GuardClient() guard.Client
	UsersClient() users.Client
}

type Dependencies struct {
	authClient  auth.Client
	db          *database.DB
	guardClient guard.Client
	usersClient users.Client
}

var _ Deps = (*Dependencies)(nil)

func (d *Dependencies) AuthClient() auth.Client {
	return d.authClient
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
