package grants

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"boilerplate/guard"
	"boilerplate/types/nullable"
	"net"
	"time"
)

const tableName = "grants"

func Create(ctx context.Context, dbc *pgxpool.Pool, userId int64,
	action guard.ActionType, foreignId int64, ip net.IP, userAgent string,
	otpId nullable.Int64, token nullable.String, validity time.Duration) (string, error) {

	var id string
	now := time.Now().UTC()

	err := dbc.QueryRow(ctx, "insert into "+tableName+" ("+
		"user_id,"+
		"action,"+
		"foreign_id,"+
		"ip,"+
		"user_agent,"+
		"otp_id,"+
		"token,"+
		"created_at,"+
		"expires_at"+
		") values ($1,$2,$3,$4,$5,$6,$7,$8,$9) "+
		"returning id",
		userId, action, foreignId, ip, userAgent,
		otpId.Value(), token.Value(), now, now.Add(validity)).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func Lookup(ctx context.Context, dbc *pgxpool.Pool,
	id string, ip net.IP) (*guard.Grant, error) {
	var g guard.Grant
	err := pgxscan.Get(ctx, dbc, &g, "select "+
		"id,"+
		"user_id,"+
		"action,"+
		"foreign_id,"+
		"ip,"+
		"user_agent,"+
		"otp_id,"+
		"token,"+
		"void,"+
		"voided_by_id,"+
		"finalised_at,"+
		"expires_at,"+
		"created_at "+
		"from "+tableName+" where "+
		"id=$1 and ip=$2", id, ip)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, guard.ErrGrantNotFound
	} else if err != nil {
		return nil, err
	}

	return &g, nil
}

func Exhaust(ctx context.Context, dbc *pgxpool.Pool, id string) error {
	_, err := dbc.Exec(ctx, "update "+tableName+" set "+
		"finalised_at=$2 where id=$1", id, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}
