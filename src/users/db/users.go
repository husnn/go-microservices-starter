package db

import (
	"boilerplate/database"
	"boilerplate/types/nullable"
	"boilerplate/users"
	"boilerplate/utils"
	"boilerplate/utils/password"
	"boilerplate/utils/random"
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/luno/jettison/errors"
	"net"
	"net/mail"
	"strings"
	"testing"
	"time"
	"unicode"
)

const (
	dbname     = "users"
	tableName  = "users"
	schemaPath = "src/users/db/schema.sql"
)

func NewTestingDB(ctx context.Context,
	t *testing.T) (*pgxpool.Pool, error) {
	return database.NewTestingDB(ctx, t, dbname,
		utils.JoinProjectPath(schemaPath))
}

func SeedInternal(ctx context.Context, dbc *pgxpool.Pool) error {
	var n int64
	err := dbc.QueryRow(ctx, "select count(*) from "+tableName).Scan(&n)
	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}

	_, err = Create(ctx, dbc, users.UserTypeInternal,
		nullable.NewString("admin@boilerplate.pk"),
		nullable.NewNull[string](), "admin123",
		net.IPv4zero, WithUserId(100))
	if err != nil {
		return err
	}
	return nil
}

func normaliseEmail(str string) (string, error) {
	parsed, err := mail.ParseAddress(str)
	if err != nil {
		return "", users.ErrInvalidEmail
	}
	return parsed.Address, nil
}

func normalisedPhone(str string) (string, error) {
	str = strings.TrimSpace(str)
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return "", users.ErrInvalidPhone
		}
	}
	return str, nil
}

func ValidPassword(str string) bool {
	return len(str) >= 8 && len(str) < 32
}

type createOpts struct {
	id int64
}

type CreateOpts func(opts *createOpts)

func WithUserId(id int64) CreateOpts {
	return func(o *createOpts) {
		o.id = id
	}
}

func Create(ctx context.Context, dbc *pgxpool.Pool, ut users.UserType,
	email nullable.String, phone nullable.String, pass string,
	signupIP net.IP, opts ...CreateOpts) (int64, error) {
	id, err := random.Int63()
	if err != nil {
		return 0, err
	}

	if !email.Null() {
		normalised, err := normaliseEmail(email.ValueOrEmpty())
		if err != nil {
			return 0, err
		}
		email = nullable.NewString(normalised)
	}

	if !phone.Null() {
		normalised, err := normalisedPhone(phone.ValueOrEmpty())
		if err != nil {
			return 0, err
		}
		phone = nullable.NewString(normalised)
	}

	if !ValidPassword(pass) {
		return 0, users.ErrInvalidPassword
	}

	pass, err = password.Hash(pass)
	if err != nil {
		return 0, err
	}

	var o createOpts
	for _, opt := range opts {
		opt(&o)
	}

	if o.id > 0 {
		id = o.id
	}

	_, err = dbc.Exec(ctx, "insert into "+tableName+
		" (id, type, email, phone, password, signup_ip, created_at) values"+
		" ($1, $2, $3, $4, $5, $6, $7)", id, ut, email.Value(),
		phone.Value(), pass, signupIP, time.Now().UTC())
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Lookup(ctx context.Context, dbc *pgxpool.Pool,
	id int64) (*users.User, error) {
	var u users.User

	err := pgxscan.Get(ctx, dbc, &u, "select * "+
		"from "+tableName+" where id=$1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, users.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func FindByEmail(ctx context.Context, dbc *pgxpool.Pool,
	ut users.UserType, email string) (*users.User, error) {
	var u users.User

	email, err := normaliseEmail(email)
	if err != nil {
		return nil, err
	}

	err = pgxscan.Get(ctx, dbc, &u, "select * "+
		"from "+tableName+" where type=$1 and email=$2", ut, email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, users.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func FindByPhone(ctx context.Context, dbc *pgxpool.Pool,
	ut users.UserType, phone string) (*users.User, error) {
	var u users.User

	phone, err := normalisedPhone(phone)
	if err != nil {
		return nil, err
	}

	err = pgxscan.Get(ctx, dbc, &u, "select * "+
		"from "+tableName+" where type=$1 and phone=$2", ut, phone)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, users.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func FindByIdentifier(ctx context.Context, dbc *pgxpool.Pool,
	ut users.UserType, identifier string) (*users.User, error) {
	email, err := normaliseEmail(identifier)
	if err != nil {
		phone, err := normalisedPhone(identifier)
		if err != nil {
			return nil, err
		}
		identifier = phone
	} else {
		identifier = email
	}

	var u users.User
	err = pgxscan.Get(ctx, dbc, &u, "select * "+
		"from "+tableName+" where type=$1 and (email=$2 or phone=$2)", ut, identifier)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, users.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}

func UpdatePassword(ctx context.Context, dbc *pgxpool.Pool,
	userID int64, pass string) (string, error) {
	if !ValidPassword(pass) {
		return "", users.ErrInvalidPassword
	}

	pass, err := password.Hash(pass)
	if err != nil {
		return "", err
	}

	_, err = dbc.Exec(ctx, "update "+tableName+" set password=$2 "+
		"where id=$1", userID, pass)
	if err != nil {
		return "", err
	}
	return pass, nil
}
