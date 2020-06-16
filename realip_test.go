package rip

import (
	"net/http"
	"testing"
)

func TestFromRequest(t *testing.T) {
	// helper for testing
	newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
		h := http.Header{}
		h.Set("X-Real-IP", xRealIP)
		for _, address := range xForwardedFor {
			h.Set("X-Forwarded-For", address)
		}

		return &http.Request{
			RemoteAddr: remoteAddr,
			Header:     h,
		}
	}

	// init test data
	publicAddr1 := "144.12.54.87"
	publicAddr2 := "119.14.55.11"
	localAddr := "127.0.0.0"

	testData := []struct {
		name     string
		request  *http.Request
		expected string
	}{
		{
			name:     "No header",
			request:  newRequest(publicAddr1, ""),
			expected: publicAddr1,
		},
		{
			name:     "Has X-Forwarded-For",
			request:  newRequest("", "", publicAddr1),
			expected: publicAddr1,
		},
		{
			name:     "Has multiple X-Forwarded-For",
			request:  newRequest("", "", localAddr, publicAddr1, publicAddr2),
			expected: publicAddr2,
		},
		{
			name:     "Has X-Real-IP",
			request:  newRequest("", publicAddr1),
			expected: publicAddr1,
		},
	}

	// Run test
	for _, v := range testData {
		if actual := FromRequest(v.request, nil); v.expected != actual {
			t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
		}
	}
}
