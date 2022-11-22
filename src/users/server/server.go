package server

import (
	"context"
	"github.com/luno/jettison/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"boilerplate/users"
	"boilerplate/users/ops"
	"boilerplate/users/state"
	"boilerplate/users/userspb"
)

var _ userspb.UsersServer = (*Server)(nil)

type Server struct {
	d state.Deps
	userspb.UnimplementedUsersServer
}

func NewServer(deps state.Deps) *Server {
	return &Server{d: deps}
}

func (s Server) Signup(ctx context.Context,
	req *userspb.SignupRequest) (*userspb.SignupResponse, error) {
	id, err := ops.Signup(ctx, s.d, users.UserType(req.Type),
		req.Email, req.Phone, req.Password, net.ParseIP(req.Ip))
	if err != nil {
		return nil, errors.Wrap(err, "error signing up user")
	}
	return &userspb.SignupResponse{
		Id: id,
	}, nil
}

func (s Server) Lookup(ctx context.Context,
	req *userspb.LookupRequest) (*userspb.User, error) {
	user, err := ops.Lookup(ctx, s.d, req.Id,
		users.UserType(req.Type), req.GetEmail(), req.GetPhone())
	if err != nil {
		return nil, errors.Wrap(err, "error getting user")
	}
	return userspb.UserToProto(user), nil
}

func (s Server) LookupForEmail(ctx context.Context,
	req *userspb.LookupRequest) (*userspb.User, error) {
	user, err := ops.Lookup(ctx, s.d, req.Id,
		users.UserType(req.Type), req.GetEmail(), req.GetPhone())
	if err != nil {
		return nil, errors.Wrap(err, "error getting user")
	}
	return userspb.UserToProto(user), nil
}

func (s Server) LookupForPhone(ctx context.Context,
	req *userspb.LookupRequest) (*userspb.User, error) {
	user, err := ops.Lookup(ctx, s.d, req.Id,
		users.UserType(req.Type), req.GetEmail(), req.GetPhone())
	if err != nil {
		return nil, errors.Wrap(err, "error getting user")
	}
	return userspb.UserToProto(user), nil
}

func (s Server) RequestPasswordReset(ctx context.Context,
	req *userspb.RequestPasswordResetRequest) (
	*userspb.RequestPasswordResetResponse, error) {
	gid, nextOtpSend, err := ops.RequestPasswordReset(ctx, s.d,
		users.UserType(req.UserType), req.Email, req.Phone,
		net.ParseIP(req.Ip), req.UserAgent)
	if err != nil {
		return nil, errors.Wrap(err, "error requesting password reset")
	}
	return &userspb.RequestPasswordResetResponse{
		GrantId:     gid,
		NextOtpSend: timestamppb.New(*nextOtpSend),
	}, nil
}

func (s Server) ResetPassword(ctx context.Context,
	req *userspb.ResetPasswordRequest) (
	*userspb.ResetPasswordResponse, error) {
	err := ops.ResetPassword(ctx, s.d, req.GrantId,
		req.Password, net.ParseIP(req.Ip), req.Otp)
	if err != nil {
		return nil, errors.Wrap(err, "error resetting password")
	}
	return &userspb.ResetPasswordResponse{}, nil
}
