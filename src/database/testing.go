package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"testing"
)

const (
	testUser = "testing"
	testPass = "testing"
)

var testDbc *pgxpool.Pool

type cleanupFunction = func(ctx context.Context, dbc *pgxpool.Pool) error

func NewTestingDB(ctx context.Context, t *testing.T, dbname,
	schemaPath string, cleanupFns ...cleanupFunction) (*pgxpool.Pool, error) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, err
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.4",
		Env: []string{
			"POSTGRES_USER=" + testUser,
			"POSTGRES_PASSWORD=" + testPass,
			"POSTGRES_DB=" + dbname,
			"listen_addresses = '*'",
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		return nil, err
	}

	port := resource.GetPort("5432/tcp")

	err = pool.Retry(func() error {
		testDbc, err = pgxpool.Connect(ctx, fmt.Sprintf(
			"postgres://%s:%s@localhost:%s/%s",
			testUser, testPass, port, dbname))
		return err
	})
	if err != nil {
		t.Logf("could not connect to database: %s", err)
		return nil, err
	}

	err = ExecuteSQLFile(ctx, testDbc, schemaPath)
	if err != nil {
		t.Fatalf("could not execute "+
			"queries from sql file: %v", err)
	}

	t.Cleanup(func() {
		if len(cleanupFns) < 1 {
			err := DropSchemas(ctx, testDbc)
			if err != nil {
				t.Logf("failed to rebuild schema: %v", err)
			}
		}

		for _, f := range cleanupFns {
			err := f(ctx, testDbc)
			if err != nil {
				t.Logf("cleanup function "+
					"failed to run successfully: %v", err)
			}
		}
		testDbc.Close()

		err = pool.Purge(resource)
		if err != nil {
			t.Logf("could not purge testing container %s%v",
				resource.Container.Name, err)
			return
		}

		t.Logf("purged container %s", resource.Container.Name)
	})

	return testDbc, nil
}
