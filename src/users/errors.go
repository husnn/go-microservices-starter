package users

import (
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
)

var (
	ErrInvalidType       = errors.New("invalid type")
	ErrInvalidEmail      = errors.New("invalid email address provided")
	ErrInvalidPhone      = errors.New("invalid phone number provided")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrUserNotFound      = errors.New("no user found for query", j.C("ERR_USER_NOT_FOUND"))
	ErrUserAlreadyExists = errors.New("user already exists")
)
