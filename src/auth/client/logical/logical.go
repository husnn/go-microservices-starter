package logical

import (
	"context"
	"net"
	"boilerplate/auth"
	"boilerplate/auth/ops"
	"boilerplate/auth/state"
	"boilerplate/users"
)

var _ auth.Client = (*Client)(nil)

type Client struct {
	d state.Deps
}

func NewClient(deps state.Deps) *Client {
	return &Client{
		d: deps,
	}
}

func (c Client) LoginWithEmail(ctx context.Context,
	ut users.UserType, email, password string, ip net.IP,
	userAgent string) (string, string, error) {
	return ops.LoginWithEmail(ctx, c.d, ut,
		email, password, ip, userAgent)
}

func (c Client) LoginWithPhone(ctx context.Context,
	ut users.UserType, phone, password string, ip net.IP,
	userAgent string) (string, string, error) {
	return ops.LoginWithPhone(ctx, c.d, ut,
		phone, password, ip, userAgent)
}

func (c Client) SubmitOTP(ctx context.Context, grantId string,
	ip net.IP, userAgent, code string) (string, error) {
	return ops.SubmitOTP(ctx, c.d, grantId, ip, userAgent, code)
}

func (c Client) LoginUnsafe(ctx context.Context, userId int64,
	ip net.IP, userAgent string) (string, error) {
	return ops.LoginUnsafe(ctx, c.d, userId, ip, userAgent)
}

func (c Client) ValidateSession(ctx context.Context,
	sessionId string, ip net.IP) (int64, error) {
	return ops.ValidateSession(ctx, c.d, sessionId, ip)
}

func (c Client) Signout(ctx context.Context, sessionId string) error {
	return ops.Signout(ctx, c.d, sessionId)
}
