package users

import (
	"net"
	"time"
	"boilerplate/types/nullable"
)

type UserType int

const (
	UserTypeUnspecified = 0
	UserTypeInternal    = 1
	UserTypeCustomer    = 2
	UserTypeSentinel    = 3
)

func (ut UserType) Valid() bool {
	return ut > UserTypeUnspecified &&
		ut < UserTypeSentinel
}

type User struct {
	Id            int64
	Type          UserType
	Email         nullable.String
	EmailVerified bool
	Phone         nullable.String
	PhoneVerified bool
	Password      string
	SignupIP      net.IP
	CreatedAt     time.Time
}

func (u User) Customer() bool {
	return u.Type == UserTypeCustomer
}

func (u User) Internal() bool {
	return u.Type == UserTypeInternal
}

type PasswordResetRequest struct {
	Id         int64
	UserId     int64
	Identifier string
	Ip         net.IP
	UserAgent  nullable.String
	CreatedAt  time.Time
}

type PasswordReset struct {
	Id          int64
	UserId      int64
	OldPassword nullable.String
	NewPassword string
	RequestId   nullable.Int64
	GrantId     nullable.String
	CreatedAt   time.Time
}
