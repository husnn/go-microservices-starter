package sessions

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"time"
	"boilerplate/auth"
	"boilerplate/types/nullable"
	"boilerplate/utils/random"
)

const tableName = "sessions"

func Create(ctx context.Context, dbc *pgxpool.Pool, userId int64,
	loginId nullable.Int64, grantId nullable.String, ip net.IP) (string, error) {

	id, err := random.Token(32)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	_, err = dbc.Exec(ctx, "insert into "+tableName+" ("+
		"id,"+
		"user_id,"+
		"login_id,"+
		"grant_id,"+
		"last_active_at,"+
		"last_active_ip,"+
		"created_at"+
		") values ($1,$2,$3,$4,$5,$6,$7)",
		id, userId, loginId.Value(), grantId.Value(), now, ip, now)
	if err != nil {
		return "", err
	}
	return id, nil
}

func Lookup(ctx context.Context, dbc *pgxpool.Pool,
	id string) (*auth.Session, error) {
	var s auth.Session
	err := pgxscan.Get(ctx, dbc, &s, "select "+
		"id,"+
		"user_id,"+
		"login_id,"+
		"grant_id,"+
		"last_active_at,"+
		"last_active_ip,"+
		"created_at,"+
		"signed_out_at "+
		"from "+tableName+" where "+
		"id=$1", id)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func UpdateLastActive(ctx context.Context, dbc *pgxpool.Pool,
	id string, ip net.IP) error {
	_, err := dbc.Exec(ctx, "update "+tableName+" set "+
		"last_active_at=$2, last_active_ip=$3 where id=$1",
		id, time.Now().UTC(), ip)
	if err != nil {
		return err
	}
	return nil
}

func Signout(ctx context.Context, dbc *pgxpool.Pool,
	id string) error {
	_, err := dbc.Exec(ctx, "update "+tableName+" set "+
		"signed_out_at=$2 "+
		"where id=$1",
		id, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}
