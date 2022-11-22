package auth

import (
	"boilerplate/gateway/httpx"
	"boilerplate/gateway/state"
	"net/http"
)

func Signout(deps state.Deps) httpx.Handler {
	return func(w http.ResponseWriter, r *httpx.Request) {
		err := deps.AuthClient().Signout(r.Context(), r.SessionId)
		if err != nil {
			httpx.Fail(w, err)
			return
		}

		httpx.DefaultSessionCookie.ClearSession(w)

		httpx.Ok(w, httpx.WithMessage("Signed out"))
	}
}
