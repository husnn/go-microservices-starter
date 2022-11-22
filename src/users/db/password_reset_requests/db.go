package password_reset_requests

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"boilerplate/types/nullable"
	"boilerplate/users"
	"net"
	"time"
)

const tableName = "users.password_reset_requests"

func Create(ctx context.Context, dbc *pgxpool.Pool,
	userId int64, identifier string, ip net.IP,
	userAgent nullable.String) (int64, error) {
	var id int64
	err := dbc.QueryRow(ctx, "insert into "+tableName+" ("+
		"user_id,"+
		"identifier,"+
		"ip,"+
		"user_agent,"+
		"created_at"+
		") values ($1,$2,$3,$4,$5) "+
		"returning id", userId, identifier,
		ip, userAgent.Value(), time.Now().UTC()).
		Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Lookup(ctx context.Context, dbc *pgxpool.Pool,
	id int64) (*users.PasswordResetRequest, error) {
	var request users.PasswordResetRequest
	err := pgxscan.Get(ctx, dbc, &request, "select * from "+
		tableName+" where id=$1", id)
	if err != nil {
		return nil, err
	}
	return &request, nil
}
