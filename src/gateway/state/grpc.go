//go:build !logical

package state

import (
	authGrpc "boilerplate/auth/client/grpc"
	guardGrpc "boilerplate/guard/client/grpc"
	usersGrpc "boilerplate/users/client/grpc"
)

func New() Deps {
	return &Dependencies{
		usersClient: usersGrpc.NewClient(),
		guardClient: guardGrpc.NewClient(),
		authClient:  authGrpc.NewClient(),
	}
}
