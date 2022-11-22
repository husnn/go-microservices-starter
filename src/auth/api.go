package auth

import (
	"context"
	"net"
	"boilerplate/users"
)

type Client interface {
	LoginWithEmail(ctx context.Context, ut users.UserType, email, password string, ip net.IP, userAgent string) (string, string, error)
	LoginWithPhone(ctx context.Context, ut users.UserType, phone, password string, ip net.IP, userAgent string) (string, string, error)
	SubmitOTP(ctx context.Context, grantId string, ip net.IP, userAgent, code string) (string, error)
	LoginUnsafe(ctx context.Context, userId int64, ip net.IP, userAgent string) (string, error)
	ValidateSession(ctx context.Context, sessionId string, ip net.IP) (int64, error)
	Signout(ctx context.Context, sessionId string) error
}
