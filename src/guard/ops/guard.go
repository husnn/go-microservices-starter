package ops

import (
	"context"
	"fmt"
	"github.com/luno/jettison/errors"
	"github.com/rs/zerolog/log"
	"net"
	"time"
	"boilerplate/env"
	"boilerplate/guard"
	"boilerplate/guard/db/grants"
	"boilerplate/guard/db/otps"
	"boilerplate/guard/state"
	"boilerplate/types/nullable"
	"boilerplate/utils/random"
)

const SendInterval = time.Minute

func Require2FA(ctx context.Context, d state.Deps,
	userID int64, action guard.ActionType, foreignID int64,
	ip net.IP, userAgent string) (string, *time.Time, error) {

	user, err := d.UsersClient().Lookup(ctx, userID)
	if err != nil {
		return "", nil, fmt.Errorf("error looking up "+
			"user for 2fa: %v", err)
	}

	channel := guard.ChannelEmail
	if user.Phone.NotNull() {
		channel = guard.ChannelSMS
	}

	code, err := random.Digits(guard.OTPLength)
	if err != nil {
		return "", nil, fmt.Errorf("error generating "+
			"random digits for otp: %v", err)
	}

	otp, err := otps.Create(ctx, d.DB().Master, user.Email,
		user.Phone, channel, code, 1, time.Now())
	if err != nil {
		return "", nil, fmt.Errorf("error creating "+
			"otp for 2fa: %v", err)
	}

	err = SendOTP(ctx, d, otp.Code)
	if err != nil {
		return "", nil, errors.Wrap(err, "error sending otp")
	}

	var validity time.Duration
	switch action {
	case guard.ActionLogin:
		validity = time.Minute * 10
	default:
		validity = time.Minute * 5
	}

	token, err := random.Token(32)
	if err != nil {
		return "", nil, fmt.Errorf("error generating "+
			"random token for grant: %v", err)
	}

	id, err := grants.Create(ctx, d.DB().Master, userID,
		action, foreignID, ip, userAgent, nullable.NewInt64(otp.Id),
		nullable.NewString(token), validity)
	if err != nil {
		return "", nil, fmt.Errorf("error creating "+
			"grant for 2fa: %v", err)
	}

	nextSend := otp.LastSentAt.Add(SendInterval)

	return id, &nextSend, nil
}

func SendOTP(ctx context.Context, d state.Deps, code string) error {
	if env.IsDev() {
		log.Info().Msgf("Your OTP is %s", code)
	}

	return nil
}

func ResendOTP(ctx context.Context, d state.Deps,
	gid string, ip net.IP) (*time.Time, error) {
	grant, err := grants.Lookup(ctx, d.DB().ReplicaOrMaster(), gid, ip)
	if err != nil {
		return nil, errors.Wrap(err, "error getting grant")
	}

	if grant.Expired() {
		return nil, guard.ErrGrantExpired
	}

	otp, err := otps.Lookup(ctx, d.DB().ReplicaOrMaster(),
		grant.OtpId.ValueOrZero())
	if err != nil {
		return nil, errors.Wrap(err, "error getting otp")
	}

	if time.Since(otp.LastSentAt) < SendInterval {
		nextSend := otp.LastSentAt.Add(SendInterval)
		return &nextSend, guard.ErrResendTimeout
	}

	err = SendOTP(ctx, d, otp.Code)
	if err != nil {
		return nil, errors.Wrap(err, "error sending otp")
	}

	err = otps.RegisterSend(ctx, d.DB().Master, otp.Id, otp.SendCount)
	if err != nil {
		return nil, errors.Wrap(err, "error registering send attempt")
	}

	nextSend := time.Now().Add(SendInterval)

	return &nextSend, nil
}

func SubmitOTP(ctx context.Context, d state.Deps, gid string,
	ip net.IP, code string) (*guard.Grant, *time.Time, error) {
	grant, err := grants.Lookup(ctx, d.DB().ReplicaOrMaster(), gid, ip)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error looking up grant")
	}

	if grant.OtpId.Null() {
		return nil, nil, guard.ErrOtpNotFound
	}

	if grant.Expired() {
		return nil, nil, guard.ErrGrantExpired
	}

	// Do void check before finalisation check.
	// This will always ensure that the latter
	// only returns false if the grant was
	// successfully granted. Otherwise, the
	// finalisation could be a result of voiding.
	if grant.Void {
		return nil, nil, guard.ErrGrantVoid
	}

	if grant.FinalisedAt != nil {
		return nil, nil, guard.ErrGrantExhausted
	}

	otp, err := otps.Lookup(ctx, d.DB().ReplicaOrMaster(),
		grant.OtpId.ValueOrZero())
	if err != nil {
		return nil, nil, errors.Wrap(err, "error getting otp")
	}

	nextSend := otp.LastSentAt.Add(SendInterval)

	// Max attempts reached
	if otp.Attempts >= 3 {
		return grant, &nextSend, guard.ErrTooManyAttempts
	}

	err = otps.RegisterAttempt(ctx, d.DB().Master,
		otp.Id, otp.Attempts)
	if err != nil {
		return grant, &nextSend, errors.Wrap(err,
			"error updating attempts for otp")
	}

	if otp.Code != code {
		return grant, &nextSend, guard.ErrIncorrectOtp
	}

	err = grants.Exhaust(ctx, d.DB().Master, grant.Id)
	if err != nil {
		return grant, &nextSend, guard.ErrGrantExhausted
	}

	// TODO: Set email/phone verified if action is verify.

	return grant, &nextSend, nil
}
