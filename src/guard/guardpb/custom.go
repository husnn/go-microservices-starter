package guardpb

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"boilerplate/guard"
	"boilerplate/types/nullable"
	"net"
	"time"
)

func GrantToProto(g *guard.Grant) *Grant {
	otpId := g.OtpId
	if otpId.Null() {
		otpId = nullable.NewInt64(0)
	}

	var finalisedAt time.Time
	if g.FinalisedAt != nil {
		finalisedAt = *g.FinalisedAt
	}

	return &Grant{
		Id:          g.Id,
		UserId:      g.UserId,
		Action:      Action(g.Action),
		ForeignId:   g.ForeignId,
		Ip:          g.Ip.String(),
		UserAgent:   g.UserAgent,
		OtpId:       g.OtpId.ValueOrZero(),
		Token:       g.Token.ValueOrEmpty(),
		Void:        g.Void,
		VoidedById:  g.VoidedById.ValueOrZero(),
		FinalisedAt: timestamppb.New(finalisedAt),
		ExpiresAt:   timestamppb.New(g.ExpiresAt),
		CreatedAt:   timestamppb.New(g.CreatedAt),
	}
}

func GrantFromProto(g *Grant) *guard.Grant {
	finalisedAt := g.FinalisedAt.AsTime()

	return &guard.Grant{
		Id:          g.Id,
		UserId:      g.UserId,
		Action:      guard.ActionType(g.Action),
		ForeignId:   g.ForeignId,
		Ip:          net.ParseIP(g.Ip),
		UserAgent:   g.UserAgent,
		OtpId:       nullable.NewInt64(g.OtpId),
		Token:       nullable.NewString(g.Token),
		Void:        g.Void,
		VoidedById:  nullable.NewInt64(g.VoidedById),
		FinalisedAt: &finalisedAt,
		ExpiresAt:   g.ExpiresAt.AsTime(),
		CreatedAt:   g.CreatedAt.AsTime(),
	}
}
