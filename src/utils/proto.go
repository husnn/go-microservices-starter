package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TimestampOrNil(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		return timestamppb.New(*t)
	}
	return nil
}

func ProtoToNullableTime(timestamp *timestamppb.Timestamp) *time.Time {
	if timestamp != nil {
		t := timestamp.AsTime()
		return &t
	}
	return nil
}
