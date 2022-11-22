package guard

import (
	"context"
	"net"
	"time"
)

type Client interface {
	Require2FA(ctx context.Context, userId int64, action ActionType, foreignId int64, ip net.IP, userAgent string) (string, *time.Time, error)
	ResendOTP(ctx context.Context, grantId string, ip net.IP) (*time.Time, error)
	SubmitOTP(ctx context.Context, grantId string, ip net.IP, code string) (*Grant, *time.Time, error)
}
