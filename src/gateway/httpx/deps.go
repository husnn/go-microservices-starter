package httpx

import (
	"boilerplate/auth"
	"boilerplate/users"
)

type Deps struct {
	authClient  auth.Client
	usersClient users.Client
}

func (d *Deps) AuthClient() auth.Client {
	return d.authClient
}

func (d *Deps) UsersClient() users.Client {
	return d.usersClient
}
