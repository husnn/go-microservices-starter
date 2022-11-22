package main

import (
	"boilerplate/database"
	guardClient "boilerplate/guard/client/grpc"
	"boilerplate/registry"
	usersDB "boilerplate/users/db"
	"boilerplate/users/server"
	"boilerplate/users/state"
	"boilerplate/users/userspb"
	"context"
	"github.com/luno/jettison/interceptors"
	"github.com/namsral/flag"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	flag.Parse()

	l, err := net.Listen("tcp",
		registry.ServiceAddress("users"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	log.Info().Msgf("Server listening at %s", l.Addr())

	ctx := context.Background()

	db, err := database.Connect(ctx, usersDB.SeedInternal)
	if err != nil {
		log.Fatal().Err(err).
			Msg("could not connect to db")
	}

	d := state.MakeDependencies(db, guardClient.NewClient())

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryServerInterceptor))
	userspb.RegisterUsersServer(s, server.NewServer(d))

	err = s.Serve(l)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
