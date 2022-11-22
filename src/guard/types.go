package guard

import (
	"boilerplate/types/nullable"
	"net"
	"time"
)

const OTPLength = 6

type ActionType int

const (
	ActionUnknown       ActionType = 0
	ActionLogin         ActionType = 1
	ActionResetPassword ActionType = 2
	ActionSentinel      ActionType = 3
)

type Grant struct {
	Id          string
	UserId      int64
	Action      ActionType
	ForeignId   int64
	Ip          net.IP
	UserAgent   string
	OtpId       nullable.Int64
	Token       nullable.String
	Void        bool
	VoidedById  nullable.Int64
	FinalisedAt *time.Time
	ExpiresAt   time.Time
	CreatedAt   time.Time
}

func (g Grant) Expired() bool {
	return time.Now().After(g.ExpiresAt)
}

type ChannelType int

const (
	ChannelUnknown  ChannelType = 0
	ChannelEmail    ChannelType = 1
	ChannelSMS      ChannelType = 2
	ChannelSentinel ChannelType = 3
)

type OTP struct {
	Id            int64
	Channel       ChannelType
	Email         nullable.String
	Phone         nullable.String
	Code          string
	SendCount     int
	LastSentAt    time.Time
	Attempts      int
	LastAttemptAt *time.Time
	CreatedAt     time.Time
}
