package auth

import (
	"github.com/luno/jettison/errors"
)

var (
	ErrIncorrectPassword = errors.New("incorrect password provided")
)
