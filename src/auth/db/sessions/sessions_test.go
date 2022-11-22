package sessions

import (
	"context"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
	"boilerplate/auth"
	"boilerplate/auth/db"
	"boilerplate/types/nullable"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	dbc, err := db.NewTestingDB(ctx, t)
	require.Nil(t, err)

	tests := []struct {
		name   string
		uid    int64
		rid    nullable.Int64
		la     nullable.Int64
		gId    nullable.String
		ip     net.IP
		errors bool
	}{
		{
			name:   "Standard",
			uid:    1,
			rid:    nullable.NewInt64(100),
			la:     nullable.NewInt64(100),
			gId:    nullable.NewNull[string](),
			ip:     net.IPv4(0, 0, 0, 0),
			errors: false,
		},
		{
			name:   "Same_User_ID",
			uid:    2,
			rid:    nullable.NewInt64(100),
			la:     nullable.NewInt64(200),
			gId:    nullable.NewNull[string](),
			ip:     net.IPv4(127, 0, 0, 1),
			errors: false,
		},
		{
			name:   "Prevent_Duplicate_Login_Attempt_ID",
			uid:    3,
			rid:    nullable.NewInt64(100),
			la:     nullable.NewInt64(200),
			gId:    nullable.NewNull[string](),
			ip:     net.IPv4(127, 0, 0, 1),
			errors: true,
		},
		{
			name:   "Missing_IP_Address",
			uid:    3,
			rid:    nullable.NewInt64(100),
			la:     nullable.NewInt64(300),
			gId:    nullable.NewNull[string](),
			ip:     nil,
			errors: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = Create(ctx, dbc, tt.uid, tt.la, tt.gId, tt.ip)
			if tt.errors {
				require.Error(t, err)
				return
			}
			require.Nil(t, err)
		})
	}
}

func TestLookup(t *testing.T) {
	ctx := context.Background()

	dbc, err := db.NewTestingDB(ctx, t)
	require.Nil(t, err)

	tests := []struct {
		uid int64
		la  nullable.Int64
		gId nullable.String
		ip  net.IP
	}{
		{
			uid: 1,
			la:  nullable.NewInt64(100),
			gId: nullable.NewNull[string](),
			ip:  net.IPv4(0, 0, 0, 0),
		},
		{
			uid: 1,
			la:  nullable.NewInt64(200),
			gId: nullable.NewString("463c2dbf-e4f7-4a2e-9c5b-e6e5307a704c"),
			ip:  net.IPv4(127, 0, 0, 1),
		},
		{
			uid: 2,
			la:  nullable.NewInt64(300),
			gId: nullable.NewNull[string](),
			ip:  net.IPv4(192, 168, 0, 1),
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			sid, err := Create(ctx, dbc, tt.uid, tt.la, tt.gId, tt.ip)
			require.Nil(t, err)

			s, err := Lookup(ctx, dbc, sid)
			require.Nil(t, err)

			require.Equal(t, &auth.Session{
				Id:           sid,
				UserId:       tt.uid,
				CreatedAt:    s.CreatedAt,
				LoginId:      tt.la,
				GrantId:      tt.gId,
				LastActiveAt: s.LastActiveAt,
				LastActiveIP: s.LastActiveIP,
				SignedOutAt:  nil,
			}, s)
		})
	}
}
