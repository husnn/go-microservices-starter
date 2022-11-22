package main

import (
	"boilerplate/database"
	"boilerplate/guard/guardpb"
	"boilerplate/guard/server"
	"boilerplate/guard/state"
	"boilerplate/registry"
	usersClient "boilerplate/users/client/grpc"
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
		registry.ServiceAddress("guard"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	log.Info().Msgf("Server listening at %s", l.Addr())

	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal().Err(err).
			Msg("could not connect to db")
	}

	d := state.MakeDependencies(db, usersClient.NewClient())

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryServerInterceptor))
	guardpb.RegisterGuardServer(s, server.NewServer(d))

	err = s.Serve(l)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
