package grpc

import (
	"context"
	"github.com/luno/jettison/interceptors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"boilerplate/auth"
	"boilerplate/auth/authpb"
	"boilerplate/registry"
	"boilerplate/users"
)

var _ auth.Client = (*Client)(nil)

type Client struct {
	conn   *grpc.ClientConn
	client authpb.AuthClient
}

func NewClient() *Client {
	conn, err := grpc.Dial(registry.ClientAddress("auth"),
		grpc.WithUnaryInterceptor(interceptors.UnaryClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).
			Msg("error connecting to auth service")
		return nil
	}

	return &Client{
		conn:   conn,
		client: authpb.NewAuthClient(conn),
	}
}

func (c Client) LoginWithEmail(ctx context.Context,
	ut users.UserType, email, password string, ip net.IP,
	userAgent string) (string, string, error) {
	res, err := c.client.Login(ctx,
		&authpb.LoginRequest{
			UserType:   int32(ut),
			Identifier: &authpb.LoginRequest_Email{Email: email},
			Password:   password,
			Ip:         ip.String(),
			UserAgent:  userAgent,
		})
	if err != nil {
		return "", "", err
	}
	return res.GetSessionId(), res.GetGrantId(), nil
}

func (c Client) LoginWithPhone(ctx context.Context,
	ut users.UserType, phone, password string, ip net.IP,
	userAgent string) (string, string, error) {
	res, err := c.client.Login(ctx,
		&authpb.LoginRequest{
			UserType:   int32(ut),
			Identifier: &authpb.LoginRequest_Phone{Phone: phone},
			Password:   password,
			Ip:         ip.String(),
			UserAgent:  userAgent,
		})
	if err != nil {
		return "", "", err
	}
	return res.GetSessionId(), res.GetGrantId(), nil
}

func (c Client) SubmitOTP(ctx context.Context, grantId string,
	ip net.IP, userAgent, code string) (string, error) {
	res, err := c.client.SubmitOTP(ctx,
		&authpb.SubmitOTPRequest{
			GrantId:   grantId,
			Ip:        ip.String(),
			UserAgent: userAgent,
			Code:      code,
		})
	if err != nil {
		return "", err
	}
	return res.SessionId, nil
}

func (c Client) LoginUnsafe(ctx context.Context,
	userID int64, ip net.IP, userAgent string) (string, error) {
	res, err := c.client.LoginUnsafe(ctx,
		&authpb.LoginUnsafeRequest{
			UserId:    userID,
			Ip:        ip.String(),
			UserAgent: userAgent,
		})
	if err != nil {
		return "", err
	}
	return res.SessionId, nil
}

func (c Client) ValidateSession(ctx context.Context,
	sessionId string, ip net.IP) (int64, error) {
	res, err := c.client.ValidateSession(ctx,
		&authpb.ValidateSessionRequest{
			SessionId: sessionId,
			Ip:        ip.String(),
		})
	if err != nil {
		return 0, err
	}
	return res.UserId, nil
}

func (c Client) Signout(ctx context.Context,
	sessionId string) error {
	_, err := c.client.Signout(ctx,
		&authpb.SignoutRequest{
			SessionId: sessionId,
		})
	if err != nil {
		return err
	}
	return nil
}
