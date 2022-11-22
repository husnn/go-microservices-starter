package password_resets

import (
	"boilerplate/types/nullable"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const tableName = "users.password_resets"

func Create(ctx context.Context, dbc *pgxpool.Pool,
	userId int64, oldPassword nullable.String, newPassword string,
	requestId nullable.Int64, grantId nullable.String) (int64, error) {
	var id int64
	err := dbc.QueryRow(ctx, "insert into "+tableName+" ("+
		"user_id,"+
		"old_password,"+
		"new_password,"+
		"request_id,"+
		"grant_id,"+
		"created_at"+
		") values ($1,$2,$3,$4,$5,$6) "+
		"returning id",
		userId, oldPassword.Value(), newPassword,
		requestId.Value(), grantId.Value(), time.Now().UTC()).
		Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
