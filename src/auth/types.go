package auth

import (
	"net"
	"time"
	"boilerplate/types/nullable"
)

type Login struct {
	Id        int64
	UserId    int64
	IP        net.IP
	UserAgent string
	CreatedAt time.Time
}

type Session struct {
	Id           string
	UserId       int64
	LoginId      nullable.Int64
	GrantId      nullable.String
	LastActiveAt time.Time
	LastActiveIP net.IP
	CreatedAt    time.Time
	SignedOutAt  *time.Time
}
