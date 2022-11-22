package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/namsral/flag"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"boilerplate/utils"
)

const defaultUri = "postgres://postgres:postgres@localhost:5432/postgres"

var masterURI = flag.String("db_uri", defaultUri, "Connection string for the master DB")
var replicaURI = flag.String("db_uri_replica", "", "Connection string for the replica DB")

type DB struct {
	Master  *pgxpool.Pool
	Replica *pgxpool.Pool
}

func (db *DB) ReplicaOrMaster() *pgxpool.Pool {
	if db.Replica != nil {
		return db.Replica
	}
	return db.Master
}

type seedFn func(ctx context.Context, dbc *pgxpool.Pool) error

func Connect(ctx context.Context, seedFns ...seedFn) (*DB, error) {
	var d DB

	conn, err := connectViaURI(ctx, *masterURI)
	if err != nil {
		return nil, err
	}
	d.Master = conn

	for _, fn := range seedFns {
		err := fn(ctx, conn)
		if err != nil {
			log.Error().Err(err).Msg("error seeding database")
		}
	}

	if *replicaURI == "" {
		return &d, nil
	}

	connReplica, err := connectViaURI(ctx, *replicaURI)
	if err != nil {
		return nil, err
	}
	d.Replica = connReplica

	return &d, nil
}

func connectViaURI(ctx context.Context,
	uri string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ExecuteSQLFile(ctx context.Context,
	dbc *pgxpool.Pool, path string) error {

	queries, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	for _, q := range strings.Split(string(queries), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		_, err := dbc.Exec(ctx, q)
		if err != nil {
			return err
		}
	}

	return nil
}

func ListServicesWithDB() ([]string, error) {
	var services []string
	err := filepath.Walk(utils.ProjectPath(),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) != ".sql" {
				return nil
			}

			if s := strings.SplitAfter(filepath.Join(path, "../.."),
				"src/"); len(s) > 1 {
				services = append(services, s[1])
			}

			return nil
		})
	return services, err
}

func DropSchemas(ctx context.Context, dbc *pgxpool.Pool) error {
	names, err := ListServicesWithDB()
	if err != nil {
		return fmt.Errorf("error listing services with a db: %v", err)
	}

	names = append(names, "public")

	for _, n := range names {
		_, err := dbc.Exec(ctx, "drop schema if exists "+n+" cascade;"+
			"create schema if not exists public;")
		if err != nil {
			return err
		}
	}
	return nil
}
