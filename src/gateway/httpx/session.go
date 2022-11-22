package httpx

import (
	"fmt"
	"boilerplate/env"
	"net/http"
	"net/url"
)

type SessionCookie struct {
	name     string
	sameSite http.SameSite
}

func NewSessionCookie(name string, sameSite http.SameSite) *SessionCookie {
	return &SessionCookie{name: name, sameSite: sameSite}
}

var DefaultSessionCookie = NewSessionCookie(
	"session", http.SameSiteDefaultMode)

func (sc *SessionCookie) Name() string {
	if env.IsStaging() {
		return fmt.Sprintf("staging_%s", sc.name)
	}
	return sc.name
}

func (sc *SessionCookie) Domain() string {
	if env.IsStaging() {
		base, err := url.Parse(env.BaseURL())
		if err != nil {
			// NoReturnErr: Set cookie on current domain.
			return ""
		}
		return "." + base.Host
	}
	return ""
}

func (sc *SessionCookie) sessionCookie(sessionID string) http.Cookie {
	secure := !env.IsDev()

	sameSite := sc.sameSite
	if !secure && sameSite == http.SameSiteNoneMode {
		sameSite = http.SameSiteDefaultMode
	}

	return http.Cookie{
		Name:     sc.Name(),
		Value:    sessionID,
		HttpOnly: true,
		Path:     "/",
		Secure:   secure,
		MaxAge:   43200, // 12 Hours
		Domain:   sc.Domain(),
		SameSite: sameSite,
	}
}

func (sc *SessionCookie) SetSession(w http.ResponseWriter, sessionID string) {
	c := sc.sessionCookie(sessionID)
	http.SetCookie(w, &c)
}

func (sc *SessionCookie) GetSession(r *http.Request) string {
	c, err := r.Cookie(sc.Name())
	if err != nil {
		return ""
	}
	return c.Value
}

func (sc *SessionCookie) ClearSession(w http.ResponseWriter) {
	c := sc.sessionCookie("")
	c.MaxAge = -1
	http.SetCookie(w, &c)
}
