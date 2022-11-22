package userspb

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"boilerplate/types/nullable"
	"boilerplate/users"
	"net"
)

func UserToProto(u *users.User) *User {
	return &User{
		Id:            u.Id,
		Type:          int32(u.Type),
		Email:         u.Email.ValueOrEmpty(),
		EmailVerified: u.EmailVerified,
		Phone:         u.Phone.ValueOrEmpty(),
		PhoneVerified: u.PhoneVerified,
		Password:      u.Password,
		SignupIp:      u.SignupIP.String(),
		CreatedAt:     timestamppb.New(u.CreatedAt),
	}
}

func UserFromProto(u *User) *users.User {
	return &users.User{
		Id:            u.Id,
		Type:          users.UserType(u.Type),
		Email:         nullable.NewString(u.Email),
		EmailVerified: u.EmailVerified,
		Phone:         nullable.NewString(u.Phone),
		PhoneVerified: u.PhoneVerified,
		Password:      u.Password,
		SignupIP:      net.ParseIP(u.SignupIp),
		CreatedAt:     u.CreatedAt.AsTime(),
	}
}
