// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.8.0
// source: guard.proto

package guardpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GuardClient is the client API for Guard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuardClient interface {
	Require2FA(ctx context.Context, in *Require2FARequest, opts ...grpc.CallOption) (*Require2FAResponse, error)
	SubmitOTP(ctx context.Context, in *SubmitOTPRequest, opts ...grpc.CallOption) (*SubmitOTPResponse, error)
	ResendOTP(ctx context.Context, in *ResendOTPRequest, opts ...grpc.CallOption) (*ResendOTPResponse, error)
}

type guardClient struct {
	cc grpc.ClientConnInterface
}

func NewGuardClient(cc grpc.ClientConnInterface) GuardClient {
	return &guardClient{cc}
}

func (c *guardClient) Require2FA(ctx context.Context, in *Require2FARequest, opts ...grpc.CallOption) (*Require2FAResponse, error) {
	out := new(Require2FAResponse)
	err := c.cc.Invoke(ctx, "/guardpb.Guard/Require2FA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guardClient) SubmitOTP(ctx context.Context, in *SubmitOTPRequest, opts ...grpc.CallOption) (*SubmitOTPResponse, error) {
	out := new(SubmitOTPResponse)
	err := c.cc.Invoke(ctx, "/guardpb.Guard/SubmitOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guardClient) ResendOTP(ctx context.Context, in *ResendOTPRequest, opts ...grpc.CallOption) (*ResendOTPResponse, error) {
	out := new(ResendOTPResponse)
	err := c.cc.Invoke(ctx, "/guardpb.Guard/ResendOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuardServer is the server API for Guard service.
// All implementations must embed UnimplementedGuardServer
// for forward compatibility
type GuardServer interface {
	Require2FA(context.Context, *Require2FARequest) (*Require2FAResponse, error)
	SubmitOTP(context.Context, *SubmitOTPRequest) (*SubmitOTPResponse, error)
	ResendOTP(context.Context, *ResendOTPRequest) (*ResendOTPResponse, error)
	mustEmbedUnimplementedGuardServer()
}

// UnimplementedGuardServer must be embedded to have forward compatible implementations.
type UnimplementedGuardServer struct {
}

func (UnimplementedGuardServer) Require2FA(context.Context, *Require2FARequest) (*Require2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Require2FA not implemented")
}
func (UnimplementedGuardServer) SubmitOTP(context.Context, *SubmitOTPRequest) (*SubmitOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOTP not implemented")
}
func (UnimplementedGuardServer) ResendOTP(context.Context, *ResendOTPRequest) (*ResendOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendOTP not implemented")
}
func (UnimplementedGuardServer) mustEmbedUnimplementedGuardServer() {}

// UnsafeGuardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuardServer will
// result in compilation errors.
type UnsafeGuardServer interface {
	mustEmbedUnimplementedGuardServer()
}

func RegisterGuardServer(s grpc.ServiceRegistrar, srv GuardServer) {
	s.RegisterService(&Guard_ServiceDesc, srv)
}

func _Guard_Require2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Require2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuardServer).Require2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guardpb.Guard/Require2FA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuardServer).Require2FA(ctx, req.(*Require2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guard_SubmitOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuardServer).SubmitOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guardpb.Guard/SubmitOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuardServer).SubmitOTP(ctx, req.(*SubmitOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guard_ResendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuardServer).ResendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guardpb.Guard/ResendOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuardServer).ResendOTP(ctx, req.(*ResendOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Guard_ServiceDesc is the grpc.ServiceDesc for Guard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Guard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "guardpb.Guard",
	HandlerType: (*GuardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Require2FA",
			Handler:    _Guard_Require2FA_Handler,
		},
		{
			MethodName: "SubmitOTP",
			Handler:    _Guard_SubmitOTP_Handler,
		},
		{
			MethodName: "ResendOTP",
			Handler:    _Guard_ResendOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guard.proto",
}