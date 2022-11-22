package grants

import (
	"context"
	"github.com/stretchr/testify/require"
	"boilerplate/guard"
	"boilerplate/guard/db"
	"boilerplate/types/nullable"
	"net"
	"testing"
	"time"
)

func TestLookupForUser(t *testing.T) {
	type create struct {
		userId    int64
		action    guard.ActionType
		foreignId int64
		ip        net.IP
		userAgent string
		OtpId     nullable.Int64
		token     nullable.String
		validity  time.Duration
	}

	tests := []struct {
		name   string
		create create
		expect *guard.Grant
		err    error
	}{
		{
			name: "",
			create: create{
				userId:    1,
				action:    guard.ActionLogin,
				foreignId: 1,
				ip:        net.IPv4(127, 0, 0, 1),
				userAgent: "UA-Browser",
				OtpId:     nullable.NewInt64(100),
				token:     nullable.NewNull[string](),
				validity:  time.Minute,
			},
			expect: &guard.Grant{
				UserId:     1,
				Action:     guard.ActionLogin,
				ForeignId:  1,
				Ip:         net.IPv4(127, 0, 0, 1).To4(),
				UserAgent:  "UA-Browser",
				OtpId:      nullable.NewInt64(100),
				Token:      nullable.NewNull[string](),
				Void:       false,
				VoidedById: nullable.NewNull[int64](),
			},
			err: nil,
		},
	}

	ctx := context.Background()
	dbc, err := db.NewTestingDB(ctx, t)
	require.Nil(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := Create(ctx, dbc,
				tt.create.userId,
				tt.create.action,
				tt.create.foreignId,
				tt.create.ip,
				tt.create.userAgent,
				tt.create.OtpId,
				tt.create.token,
				tt.create.validity)
			require.Nil(t, err)

			created, err := Lookup(ctx, dbc, id, tt.create.ip)
			require.Nil(t, err)

			tt.expect.Id = id
			tt.expect.ExpiresAt = created.CreatedAt.Add(tt.create.validity)
			tt.expect.CreatedAt = created.CreatedAt

			require.Equal(t, tt.expect, created)
		})
	}
}
