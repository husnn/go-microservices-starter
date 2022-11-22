package logical

import (
	"context"
	"net"
	"time"
	"boilerplate/users"
	"boilerplate/users/ops"
	"boilerplate/users/state"
)

var _ users.Client = (*Client)(nil)

type Client struct {
	d state.Deps
}

func NewClient(deps state.Deps) *Client {
	return &Client{
		d: deps,
	}
}

func (c Client) Signup(ctx context.Context, ut users.UserType,
	email string, phone string, password string, ip net.IP) (int64, error) {
	return ops.Signup(ctx, c.d, ut, email, phone, password, ip)
}

func (c Client) Lookup(ctx context.Context,
	id int64) (*users.User, error) {
	return ops.Lookup(ctx, c.d, id,
		users.UserTypeUnspecified, "", "")
}

func (c Client) LookupForEmail(ctx context.Context,
	ut users.UserType, email string) (*users.User, error) {
	return ops.Lookup(ctx, c.d, 0, ut, email, "")
}

func (c Client) LookupForPhone(ctx context.Context,
	ut users.UserType, phone string) (*users.User, error) {
	return ops.Lookup(ctx, c.d, 0, ut, "", phone)
}

func (c Client) RequestPasswordReset(ctx context.Context, ut users.UserType,
	email, phone string, ip net.IP, userAgent string) (string, *time.Time, error) {
	return ops.RequestPasswordReset(ctx, c.d, ut, email, phone, ip, userAgent)
}

func (c Client) ResetPassword(ctx context.Context,
	grantId, password string, ip net.IP, otp string) error {
	return ops.ResetPassword(ctx, c.d, grantId, password, ip, otp)
}
