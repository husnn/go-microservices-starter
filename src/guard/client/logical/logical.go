package logical

import (
	"context"
	"net"
	"time"
	"boilerplate/guard"
	"boilerplate/guard/ops"
	"boilerplate/guard/state"
)

var _ guard.Client = (*Client)(nil)

type Client struct {
	d state.Deps
}

func NewClient(deps state.Deps) *Client {
	return &Client{
		d: deps,
	}
}

func (c Client) Require2FA(ctx context.Context, userId int64, action guard.ActionType,
	foreignId int64, ip net.IP, userAgent string) (string, *time.Time, error) {
	return ops.Require2FA(ctx, c.d, userId, action, foreignId, ip, userAgent)
}

func (c Client) ResendOTP(ctx context.Context, grantId string,
	ip net.IP) (*time.Time, error) {
	return ops.ResendOTP(ctx, c.d, grantId, ip)
}

func (c Client) SubmitOTP(ctx context.Context, grantId string,
	ip net.IP, code string) (*guard.Grant, *time.Time, error) {
	return ops.SubmitOTP(ctx, c.d, grantId, ip, code)
}
