package ops

import (
	"context"
	"github.com/luno/jettison/errors"
	"net"
	"time"
	"boilerplate/guard"
	"boilerplate/types/nullable"
	"boilerplate/users"
	"boilerplate/users/db"
	"boilerplate/users/db/password_reset_requests"
	"boilerplate/users/db/password_resets"
	"boilerplate/users/state"
)

func Signup(ctx context.Context, d state.Deps, ut users.UserType,
	email, phone, password string, ip net.IP) (int64, error) {
	if !ut.Valid() {
		return 0, users.ErrInvalidType
	}

	if email == "" && phone == "" {
		return 0, errors.New("missing email and phone number")
	}

	var existing *users.User
	var err error

	if email != "" {
		existing, err = db.FindByEmail(ctx, d.DB().ReplicaOrMaster(), ut, email)
		if err != nil && !errors.Is(err, users.ErrUserNotFound) {
			return 0, errors.Wrap(err, "error checking user for email")
		}
	}

	if existing == nil {
		existing, err = db.FindByPhone(ctx, d.DB().ReplicaOrMaster(), ut, phone)
		if err != nil && !errors.Is(err, users.ErrUserNotFound) {
			return 0, errors.Wrap(err, "error checking user for phone")
		}
	}

	if existing != nil {
		return 0, users.ErrUserAlreadyExists
	}

	return db.Create(ctx, d.DB().Master, ut, nullable.NewString(email),
		nullable.NewString(phone), password, ip)
}

func Lookup(ctx context.Context, d state.Deps, id int64,
	ut users.UserType, email, phone string) (*users.User, error) {
	if id > 0 {
		return db.Lookup(ctx, d.DB().ReplicaOrMaster(), id)
	}
	if len(email) > 0 {
		return db.FindByEmail(ctx, d.DB().ReplicaOrMaster(), ut, email)
	}
	if len(phone) > 0 {
		return db.FindByPhone(ctx, d.DB().ReplicaOrMaster(), ut, phone)
	}
	return nil, errors.New("no valid identifier")
}

func RequestPasswordReset(ctx context.Context, d state.Deps, ut users.UserType,
	email, phone string, ip net.IP, ua string) (string, *time.Time, error) {
	var user *users.User
	var err error
	var identifier string

	if email != "" {
		user, err = db.FindByEmail(ctx, d.DB().ReplicaOrMaster(), ut, email)
		identifier = email
	} else {
		user, err = db.FindByPhone(ctx, d.DB().ReplicaOrMaster(), ut, phone)
		identifier = phone
	}
	if err != nil {
		return "", nil, errors.Wrap(err, "error looking up user")
	}

	requestId, err := password_reset_requests.Create(ctx, d.DB().Master,
		user.Id, identifier, ip, nullable.NewString(ua))
	if err != nil {
		return "", nil, errors.Wrap(err, "error creating reset request")
	}

	gid, nextOtpSend, err := d.GuardClient().Require2FA(ctx, user.Id,
		guard.ActionResetPassword, requestId, ip, ua)
	if err != nil {
		return "", nil, errors.Wrap(err, "error creating grant")
	}

	return gid, nextOtpSend, nil
}

func ResetPassword(ctx context.Context, d state.Deps,
	gid, password string, ip net.IP, otp string) error {
	if !db.ValidPassword(password) {
		return users.ErrInvalidPassword
	}

	grant, _, err := d.GuardClient().SubmitOTP(ctx, gid, ip, otp)
	if err != nil {
		return errors.Wrap(err, "error submitting otp")
	}

	resetRequest, err := password_reset_requests.Lookup(ctx,
		d.DB().ReplicaOrMaster(), grant.ForeignId)
	if err != nil {
		return errors.Wrap(err, "error getting reset request")
	}

	user, err := db.Lookup(ctx, d.DB().ReplicaOrMaster(), grant.UserId)
	if errors.Is(err, users.ErrUserNotFound) {
		return err
	} else if err != nil {
		return errors.Wrap(err, "error getting user")
	}

	newPassword, err := db.UpdatePassword(ctx,
		d.DB().Master, user.Id, password)
	if errors.Is(err, users.ErrInvalidPassword) {
		return err
	} else if err != nil {
		return errors.Wrap(err, "error updating password")
	}

	_, err = password_resets.Create(ctx, d.DB().Master,
		user.Id, nullable.NewString(user.Password), newPassword,
		nullable.NewInt64(resetRequest.Id), nullable.NewString(grant.Id))
	if err != nil {
		return errors.Wrap(err, "error creating reset record")
	}

	return nil
}
