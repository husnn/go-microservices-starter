package grpc

import (
	"context"
	"github.com/luno/jettison/interceptors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
	"boilerplate/guard"
	"boilerplate/guard/guardpb"
	"boilerplate/registry"
)

var _ guard.Client = (*Client)(nil)

type Client struct {
	conn   *grpc.ClientConn
	client guardpb.GuardClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(registry.ClientAddress("guard"),
		grpc.WithUnaryInterceptor(interceptors.UnaryClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).
			Msg("error connecting to guard service")
		return nil
	}

	return &Client{
		conn:   conn,
		client: guardpb.NewGuardClient(conn),
	}
}

func (c Client) Require2FA(ctx context.Context, userId int64, action guard.ActionType,
	foreignId int64, ip net.IP, userAgent string) (string, *time.Time, error) {
	res, err := c.client.Require2FA(ctx, &guardpb.Require2FARequest{
		UserId:    userId,
		Action:    guardpb.Action(action),
		ForeignId: foreignId,
		Ip:        ip.String(),
		UserAgent: userAgent,
	})
	if err != nil {
		return "", nil, err
	}
	nextSend := res.NextSend.AsTime()
	return res.GrantId, &nextSend, nil
}

func (c Client) ResendOTP(ctx context.Context, grantId string,
	ip net.IP) (*time.Time, error) {
	res, err := c.client.ResendOTP(ctx, &guardpb.ResendOTPRequest{
		GrantId: grantId,
		Ip:      ip.String(),
	})
	if err != nil {
		return nil, err
	}
	nextSend := res.NextSend.AsTime()
	return &nextSend, nil
}

func (c Client) SubmitOTP(ctx context.Context, grantId string,
	ip net.IP, code string) (*guard.Grant, *time.Time, error) {
	res, err := c.client.SubmitOTP(ctx, &guardpb.SubmitOTPRequest{
		GrantId: grantId,
		Ip:      ip.String(),
		Code:    code,
	})
	if err != nil && res == nil {
		return nil, nil, err
	}
	nextSend := res.NextSend.AsTime()
	return guardpb.GrantFromProto(res.Grant), &nextSend, nil
}
