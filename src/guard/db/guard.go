package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
	"boilerplate/database"
	"boilerplate/utils"
)

const dbname = "guard"
const schemaPath = "src/guard/db/schema.sql"

func NewTestingDB(ctx context.Context,
	t *testing.T) (*pgxpool.Pool, error) {
	return database.NewTestingDB(ctx, t, dbname,
		utils.JoinProjectPath(schemaPath))
}
