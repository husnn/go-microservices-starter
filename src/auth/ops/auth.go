package ops

import (
	"context"
	"github.com/luno/jettison/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"net"
	"boilerplate/auth"
	"boilerplate/auth/db/logins"
	"boilerplate/auth/db/sessions"
	"boilerplate/auth/state"
	"boilerplate/guard"
	"boilerplate/types/nullable"
	"boilerplate/users"
	"boilerplate/utils/password"
)

func LoginWithEmail(ctx context.Context, d state.Deps, ut users.UserType,
	email, password string, ip net.IP, ua string) (string, string, error) {
	user, err := d.UsersClient().LookupForEmail(ctx, ut, email)
	if err != nil {
		return "", "", errors.Wrap(err, "error getting user")
	}
	sid, gid, err := loginUsingPassword(ctx, d, user, password, ip, ua)
	if err != nil {
		return "", "", errors.Wrap(err, "login error")
	}
	return sid, gid, nil
}

func LoginWithPhone(ctx context.Context, d state.Deps, ut users.UserType,
	phone, password string, ip net.IP, ua string) (string, string, error) {
	user, err := d.UsersClient().LookupForPhone(ctx, ut, phone)
	if err != nil {
		return "", "", errors.Wrap(err, "error getting user")
	}
	sid, gid, err := loginUsingPassword(ctx, d, user, password, ip, ua)
	if err != nil {
		return "", "", errors.Wrap(err, "login error")
	}
	return sid, gid, nil
}

func createSession(ctx context.Context, d state.Deps, userId int64,
	loginId int64, grantId nullable.String, ip net.IP) (string, error) {
	sid, err := sessions.Create(ctx, d.DB().Master,
		userId, nullable.NewInt64(loginId), grantId, ip)
	if err != nil {
		return "", errors.Wrap(err,
			"error creating session")
	}
	return sid, nil
}

func loginUsingPassword(ctx context.Context, d state.Deps, user *users.User,
	pwd string, ip net.IP, ua string) (string, string, error) {
	if user.Password == "" {
		return "", "", errors.New(
			"password not yet set for user")
	}

	err := password.Verify(user.Password, pwd)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", "", auth.ErrIncorrectPassword
	} else if err != nil {
		return "", "", errors.Wrap(err,
			"error comparing passwords")
	}

	loginId, err := logins.Create(ctx,
		d.DB().Master, user.Id, ip, ua)
	if err != nil {
		return "", "", errors.Wrap(err,
			"error creating login")
	}

	if user.EmailVerified || user.PhoneVerified {
		gid, _, err := d.GuardClient().Require2FA(ctx,
			user.Id, guard.ActionLogin, loginId, ip, ua)
		if err != nil {
			return "", "", errors.Wrap(err,
				"error creating grant for login")
		}
		return "", gid, nil
	}

	sid, err := createSession(ctx, d, user.Id,
		loginId, nullable.NewNull[string](), ip)
	if err != nil {
		return "", "", errors.Wrap(err,
			"could not create session")
	}
	return sid, "", nil
}

func SubmitOTP(ctx context.Context, d state.Deps, gid string,
	ip net.IP, ua, code string) (string, error) {
	grant, _, err := d.GuardClient().
		SubmitOTP(ctx, gid, ip, code)
	if err != nil {
		return "", errors.Wrap(err,
			"error submitting otp")
	}

	sid, err := createSession(ctx, d, grant.UserId,
		grant.ForeignId, nullable.NewString(gid), ip)
	if err != nil {
		return "", err
	}

	return sid, err
}

func LoginUnsafe(ctx context.Context, d state.Deps,
	uid int64, ip net.IP, ua string) (string, error) {
	loginId, err := logins.Create(ctx,
		d.DB().Master, uid, ip, ua)
	if err != nil {
		return "", errors.Wrap(err,
			"error creating login")
	}

	sid, err := sessions.Create(ctx, d.DB().Master, uid,
		nullable.NewInt64(loginId), nullable.NewNull[string](), ip)
	if err != nil {
		return "", err
	}

	return sid, nil
}

func ValidateSession(ctx context.Context, d state.Deps,
	sid string, ip net.IP) (int64, error) {
	sess, err := sessions.Lookup(ctx,
		d.DB().ReplicaOrMaster(), sid)
	if err != nil {
		return 0, errors.Wrap(err,
			"error getting session")
	}

	if sess.SignedOutAt != nil {
		return 0, errors.Wrap(err,
			"session signed out")
	}

	updateLastActive(ctx, d, sess.Id, ip)
	return sess.UserId, nil
}

func updateLastActive(ctx context.Context,
	d state.Deps, sessionId string, ip net.IP) {
	err := sessions.UpdateLastActive(ctx,
		d.DB().Master, sessionId, ip)
	if err != nil {
		log.Error().
			Msg("error updating session last active")
	}
}

func Signout(ctx context.Context,
	d state.Deps, sid string) error {
	err := sessions.Signout(ctx,
		d.DB().Master, sid)
	if err != nil {
		return errors.Wrap(err,
			"signout error")
	}
	return nil
}
