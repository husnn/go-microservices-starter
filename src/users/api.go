package users

import (
	"context"
	"net"
	"time"
)

type Client interface {
	Signup(ctx context.Context, ut UserType, email string, phone string, password string, ip net.IP) (int64, error)
	Lookup(ctx context.Context, id int64) (*User, error)
	LookupForEmail(ctx context.Context, ut UserType, email string) (*User, error)
	LookupForPhone(ctx context.Context, ut UserType, phone string) (*User, error)
	RequestPasswordReset(ctx context.Context, ut UserType, email, phone string, ip net.IP, userAgent string) (string, *time.Time, error)
	ResetPassword(ctx context.Context, grantId, password string, ip net.IP, otp string) error
}
