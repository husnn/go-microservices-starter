package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"boilerplate/gateway/state"
)

type Handler func(http.ResponseWriter, *Request)

func Base(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(w, newRequestWithMetadata(r))
	})
}

type checkFn func(r *Request) (bool, error)

type authOpts struct {
	check checkFn
}

type AuthOpts func(opts *authOpts)

func WithAdditionalCheck(fn checkFn) AuthOpts {
	return func(o *authOpts) {
		o.check = fn
	}
}

func extractAuthTokenFromHeader(r *http.Request) string {
	splitHeader := strings.Split(
		r.Header.Get("Authorization"), "Bearer ")
	if len(splitHeader) > 1 {
		return splitHeader[1]
	}
	return ""
}

func Authenticated(d state.Deps, h Handler, opts ...AuthOpts) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := newRequestWithMetadata(r)

		sid := DefaultSessionCookie.GetSession(r)

		if sid == "" {
			sid = extractAuthTokenFromHeader(r)
		}

		if sid == "" {
			Fail(w, errors.New("could not get session cookie"),
				WithMessage("Unauthenticated"), WithCode(403))
			return
		}

		userId, err := d.AuthClient().ValidateSession(r.Context(), sid, req.IP)
		if err != nil {
			Fail(w, fmt.Errorf("could not validate auth status "+
				"from session cookie: %v", err))
			return
		}

		if userId < 1 {
			Fail(w, errors.New("invalid authenticated user id"),
				WithMessage("Unauthenticated"), WithCode(403))
			return
		}

		setAuthMetadata(req, &authMetadata{
			userId:    userId,
			sessionId: sid,
		})

		var o authOpts
		for _, opt := range opts {
			opt(&o)
		}

		if o.check != nil {
			pass, err := o.check(req)
			if err != nil {
				Fail(w, fmt.Errorf("failed "+
					"additional auth check: %v", err))
				return
			}
			if !pass {
				Fail(w, nil, WithMessage("Access denied"),
					WithCode(403))
				return
			}
		}

		h(w, req)
	})
}

func Internal(d state.Deps, h Handler) http.Handler {
	return Authenticated(d, h, WithAdditionalCheck(
		func(r *Request) (bool, error) {
			user, err := d.UsersClient().
				Lookup(r.Context(), r.UserId)
			if err != nil {
				return false, err
			}

			return user.Internal(), nil
		}))
}
