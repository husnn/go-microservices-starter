package grpc

import (
	"context"
	"github.com/luno/jettison/interceptors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
	"boilerplate/registry"
	"boilerplate/users"
	"boilerplate/users/userspb"
)

var _ users.Client = (*Client)(nil)

type Client struct {
	conn   *grpc.ClientConn
	client userspb.UsersClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(registry.ClientAddress("users"),
		grpc.WithUnaryInterceptor(interceptors.UnaryClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).
			Msg("error connecting to users service")
		return nil
	}

	return &Client{
		conn:   conn,
		client: userspb.NewUsersClient(conn),
	}
}

func (c Client) Signup(ctx context.Context, ut users.UserType,
	email string, phone string, password string, ip net.IP) (int64, error) {
	res, err := c.client.Signup(ctx, &userspb.SignupRequest{
		Type:     int32(ut),
		Email:    email,
		Phone:    phone,
		Password: password,
		Ip:       ip.String(),
	})
	if err != nil {
		return 0, err
	}
	return res.Id, nil
}

func (c Client) Lookup(ctx context.Context,
	id int64) (*users.User, error) {
	res, err := c.client.Lookup(ctx, &userspb.LookupRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return userspb.UserFromProto(res), nil
}

func (c Client) LookupForEmail(ctx context.Context,
	ut users.UserType, email string) (*users.User, error) {
	res, err := c.client.LookupForEmail(ctx,
		&userspb.LookupRequest{Type: int32(ut),
			Identifier: &userspb.LookupRequest_Email{Email: email}})
	if err != nil {
		return nil, err
	}
	return userspb.UserFromProto(res), nil
}

func (c Client) LookupForPhone(ctx context.Context,
	ut users.UserType, phone string) (*users.User, error) {
	res, err := c.client.LookupForPhone(ctx,
		&userspb.LookupRequest{Type: int32(ut),
			Identifier: &userspb.LookupRequest_Phone{Phone: phone}})
	if err != nil {
		return nil, err
	}
	return userspb.UserFromProto(res), nil
}

func (c Client) RequestPasswordReset(ctx context.Context, ut users.UserType,
	email, phone string, ip net.IP, userAgent string) (string, *time.Time, error) {
	res, err := c.client.RequestPasswordReset(ctx, &userspb.RequestPasswordResetRequest{
		UserType:  int32(ut),
		Email:     email,
		Phone:     phone,
		Ip:        ip.String(),
		UserAgent: userAgent,
	})
	if err != nil {
		return "", nil, err
	}
	nextOtpSend := res.NextOtpSend.AsTime()
	return res.GrantId, &nextOtpSend, nil
}

func (c Client) ResetPassword(ctx context.Context,
	grantId, password string, ip net.IP, otp string) error {
	_, err := c.client.ResetPassword(ctx, &userspb.ResetPasswordRequest{
		GrantId:  grantId,
		Password: password,
		Ip:       ip.String(),
		Otp:      otp,
	})
	return err
}
