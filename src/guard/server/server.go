package server

import (
	"context"
	"github.com/luno/jettison/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"boilerplate/guard"
	"boilerplate/guard/guardpb"
	"boilerplate/guard/ops"
	"boilerplate/guard/state"
)

var _ guardpb.GuardServer = (*Server)(nil)

type Server struct {
	d state.Deps
	guardpb.UnimplementedGuardServer
}

func NewServer(deps state.Deps) *Server {
	return &Server{d: deps}
}

func (s Server) Require2FA(ctx context.Context,
	req *guardpb.Require2FARequest) (
	*guardpb.Require2FAResponse, error) {
	grantId, nextSend, err := ops.Require2FA(ctx, s.d,
		req.UserId, guard.ActionType(req.Action), req.ForeignId,
		net.ParseIP(req.Ip), req.UserAgent)
	if err != nil {
		return nil, err
	}
	return &guardpb.Require2FAResponse{
		GrantId:  grantId,
		NextSend: timestamppb.New(*nextSend),
	}, nil
}

func (s Server) ResendOTP(ctx context.Context,
	req *guardpb.ResendOTPRequest) (*guardpb.ResendOTPResponse, error) {
	nextSend, err := ops.ResendOTP(ctx, s.d, req.GrantId, net.ParseIP(req.Ip))
	if errors.Is(err, guard.ErrResendTimeout) {
		return &guardpb.ResendOTPResponse{
			NextSend: timestamppb.New(*nextSend)}, err
	} else if err != nil {
		return nil, errors.Wrap(err, "error sending out otp")
	}
	return &guardpb.ResendOTPResponse{
		NextSend: timestamppb.New(*nextSend),
	}, nil
}

func (s Server) SubmitOTP(ctx context.Context,
	req *guardpb.SubmitOTPRequest) (
	*guardpb.SubmitOTPResponse, error) {
	grant, nextSend, err := ops.SubmitOTP(ctx, s.d,
		req.GrantId, net.ParseIP(req.Ip), req.Code)
	if err != nil {
		err = errors.Wrap(err, "error submitting otp")
		if grant == nil || nextSend == nil {
			return nil, err
		}
	}
	return &guardpb.SubmitOTPResponse{
		Grant:    guardpb.GrantToProto(grant),
		NextSend: timestamppb.New(*nextSend),
	}, err
}
