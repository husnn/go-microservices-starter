package logins

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"time"
	"boilerplate/auth"
)

const tableName = "logins"

func Create(ctx context.Context, dbc *pgxpool.Pool,
	userId int64, ip net.IP, userAgent string) (int64, error) {
	var id int64
	err := dbc.QueryRow(ctx, "insert into "+tableName+" ("+
		"user_id,"+
		"ip,"+
		"user_agent,"+
		"created_at"+
		") values ($1,$2,$3,$4) "+
		"returning id",
		userId, ip, userAgent, time.Now().UTC()).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Lookup(ctx context.Context, dbc *pgxpool.Pool,
	id int64) (*auth.Login, error) {
	var l auth.Login
	err := pgxscan.Get(ctx, dbc, &l, "select "+
		"id,"+
		"user_id,"+
		"ip,"+
		"user_agent,"+
		"created_at "+
		"from "+tableName+" where "+
		"id=$1", id)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
