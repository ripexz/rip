package rip

import (
	"net"
	"net/http"
	"strings"
)

// FromRequest returns client's real public IP address from http request headers.
func FromRequest(r *http.Request, filter Filter) string {
	if filter == nil {
		filter = FilterDefault
	}

	// Fetch header value
	xRealIP := r.Header.Get("X-Real-Ip")
	xForwardedFor := r.Header.Get("X-Forwarded-For")

	// If both empty, return IP from remote address
	if xRealIP == "" && xForwardedFor == "" {
		var remoteIP string

		// If there are colon in remote address, remove the port number
		// otherwise, return remote address as is
		if strings.ContainsRune(r.RemoteAddr, ':') {
			remoteIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		} else {
			remoteIP = r.RemoteAddr
		}

		return remoteIP
	}

	// Check list of IP in X-Forwarded-For and return the first global address
	if address, ok := filter(strings.Split(xForwardedFor, ",")); ok {
		return address
	}

	// If nothing succeed, return X-Real-IP
	return xRealIP
}
