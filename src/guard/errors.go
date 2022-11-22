package guard

import (
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
)

var (
	ErrGrantExpired    = errors.New("the associated grant has expired")
	ErrGrantVoid       = errors.New("grant has been voided")
	ErrGrantExhausted  = errors.New("grant has already been exhausted")
	ErrGrantNotFound   = errors.New("no grant found", j.C("ERR_GRANT_NOT_FOUND"))
	ErrOtpNotFound     = errors.New("no otp associated with grant")
	ErrIncorrectOtp    = errors.New("incorrect OTP provided")
	ErrTooManyAttempts = errors.New("too many attempts")
	ErrResendTimeout   = errors.New("not enough time elapsed since last send")
)
