package otps

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"boilerplate/guard"
	"boilerplate/types/nullable"
	"time"
)

const tableName = "otps"

func Create(ctx context.Context, dbc *pgxpool.Pool,
	email, phone nullable.String, channel guard.ChannelType,
	code string, sendCount int, lastSent time.Time) (*guard.OTP, error) {
	var id int64
	err := dbc.QueryRow(ctx, "insert into "+tableName+" ("+
		"channel,"+
		"email,"+
		"phone,"+
		"code,"+
		"send_count,"+
		"last_sent_at,"+
		"created_at"+
		") values ($1, $2, $3, $4, $5, $6, $7) "+
		"returning id",
		channel,
		email.Value(),
		phone.Value(),
		code,
		sendCount,
		lastSent.UTC(),
		time.Now().UTC()).
		Scan(&id)
	if err != nil {
		return nil, err
	}

	var otp guard.OTP
	err = pgxscan.Get(ctx, dbc, &otp, "select * from "+
		tableName+" where id=$1", id)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func Lookup(ctx context.Context, dbc *pgxpool.Pool,
	id int64) (*guard.OTP, error) {
	var otp guard.OTP

	err := pgxscan.Get(ctx, dbc, &otp, "select "+
		"id,"+
		"channel,"+
		"email,"+
		"phone,"+
		"code,"+
		"send_count,"+
		"last_sent_at,"+
		"attempts,"+
		"last_attempt_at,"+
		"created_at "+
		"from "+tableName+" where id=$1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, guard.ErrOtpNotFound
	} else if err != nil {
		return nil, err
	}

	return &otp, nil
}

func RegisterAttempt(ctx context.Context, dbc *pgxpool.Pool,
	id int64, prevCount int) error {
	_, err := dbc.Exec(ctx, "update "+tableName+" set "+
		"attempts=$2, last_attempt_at=$3 "+
		"where id=$1", id, prevCount+1, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}

func RegisterSend(ctx context.Context, dbc *pgxpool.Pool,
	id int64, prevCount int) error {
	_, err := dbc.Exec(ctx, "update "+tableName+" set "+
		"send_count=$2, last_sent_at=$3 "+
		"where id=$1", id, prevCount+1, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}
