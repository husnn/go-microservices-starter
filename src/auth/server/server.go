package server

import (
	"context"
	"github.com/luno/jettison/errors"
	"net"
	"boilerplate/auth/authpb"
	"boilerplate/auth/ops"
	"boilerplate/auth/state"
	"boilerplate/users"
)

var _ authpb.AuthServer = (*Server)(nil)

type Server struct {
	d state.Deps
	authpb.UnimplementedAuthServer
}

func NewServer(deps state.Deps) *Server {
	return &Server{d: deps}
}

func (s Server) Login(ctx context.Context,
	req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	var sid, gid string
	var err error
	if len(req.GetPhone()) > 0 {
		sid, gid, err = ops.LoginWithPhone(ctx, s.d, users.UserType(req.UserType),
			req.GetPhone(), req.Password, net.ParseIP(req.Ip), req.UserAgent)
	} else if len(req.GetEmail()) > 0 {
		sid, gid, err = ops.LoginWithEmail(ctx, s.d, users.UserType(req.UserType),
			req.GetEmail(), req.Password, net.ParseIP(req.Ip), req.UserAgent)
	} else {
		return nil, errors.New("missing identifier")
	}
	if err != nil {
		return nil, err
	}

	if len(sid) > 0 {
		return &authpb.LoginResponse{
			Id: &authpb.LoginResponse_SessionId{SessionId: sid},
		}, nil
	}

	return &authpb.LoginResponse{
		Id: &authpb.LoginResponse_GrantId{GrantId: gid},
	}, nil
}

func (s Server) SubmitOTP(ctx context.Context,
	req *authpb.SubmitOTPRequest) (*authpb.SubmitOTPResponse, error) {
	sid, err := ops.SubmitOTP(ctx, s.d, req.GrantId,
		net.ParseIP(req.Ip), req.UserAgent, req.Code)
	if err != nil {
		return nil, err
	}
	return &authpb.SubmitOTPResponse{
		SessionId: sid,
	}, nil
}

func (s Server) LoginUnsafe(ctx context.Context,
	req *authpb.LoginUnsafeRequest) (*authpb.LoginUnsafeResponse, error) {
	sid, err := ops.LoginUnsafe(ctx, s.d, req.UserId,
		net.ParseIP(req.Ip), req.UserAgent)
	if err != nil {
		return nil, err
	}
	return &authpb.LoginUnsafeResponse{
		SessionId: sid,
	}, nil
}

func (s Server) ValidateSession(ctx context.Context,
	req *authpb.ValidateSessionRequest) (*authpb.ValidateSessionResponse, error) {
	uid, err := ops.ValidateSession(ctx, s.d,
		req.SessionId, net.ParseIP(req.Ip))
	if err != nil {
		return nil, err
	}
	return &authpb.ValidateSessionResponse{
		UserId: uid,
	}, nil
}

func (s Server) Signout(ctx context.Context,
	req *authpb.SignoutRequest) (*authpb.SignoutResponse, error) {
	err := ops.Signout(ctx, s.d, req.SessionId)
	if err != nil {
		return nil, err
	}
	return &authpb.SignoutResponse{}, nil
}
