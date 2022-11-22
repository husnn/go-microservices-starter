package httpx

import (
	"context"
	"golang.org/x/text/language"
	"net"
	"net/http"
	"strings"
)

const (
	headerXRealIP       = "X-Real-IP"
	headerXForwardedFor = "X-Forwarded-For"
)

const userIDContextKey = "user_id"

type authMetadata struct {
	userId    int64
	sessionId string
}

func newRequestWithMetadata(r *http.Request) *Request {
	req := &Request{
		Request: r,
		Lang:    language.English,
		IP:      GetRemoteIP(r),
	}

	return req
}

func setAuthMetadata(r *Request, auth *authMetadata) {
	if auth == nil {
		return
	}

	ctx := r.Context()

	// Add to context
	ctx = context.WithValue(ctx, userIDContextKey, auth.userId)

	// Add to request
	r.UserId = auth.userId
	r.SessionId = auth.sessionId

	r.Request = r.WithContext(ctx)
}

func GetRemoteIP(r *http.Request) net.IP {
	ip := r.Header.Get(headerXRealIP)
	if ip == "" {
		ip = r.Header.Get(headerXForwardedFor)
	}
	if ip == "" {
		ip = r.RemoteAddr
	}

	return net.ParseIP(stripPort(ip))
}

func stripPort(addr string) string {
	if strings.Contains(addr, ".") &&
		strings.Contains(addr, ":") {
		host, _, _ := net.SplitHostPort(addr)
		return host
	}
	return addr
}
