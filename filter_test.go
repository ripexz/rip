package rip

import (
	"testing"
)

func TestFilterPublicAddress(t *testing.T) {
	testData := []struct {
		ips      []string
		expected string
	}{
		{[]string{"127.0.0.0"}, ""},
		{[]string{"10.0.0.0"}, ""},
		{[]string{"169.254.0.0"}, ""},
		{[]string{"192.168.0.0"}, ""},
		{[]string{"::1"}, ""},
		{[]string{"fc00::"}, ""},

		{[]string{"172.15.0.0"}, "172.15.0.0"},
		{[]string{"172.16.0.0"}, ""},
		{[]string{"172.31.0.0"}, ""},
		{[]string{"172.32.0.0"}, "172.32.0.0"},

		{[]string{"147.12.56.11"}, "147.12.56.11"},
	}

	for _, tt := range testData {
		addr, _ := FilterPublicAddress(tt.ips)
		if addr != tt.expected {
			t.Errorf("unexpected result, want: '%s', got: '%s'", tt.expected, addr)
		}
	}
}

func TestFilterAWS(t *testing.T) {
	testData := []struct {
		ips      []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"127.0.0.1"}, "127.0.0.1"},
		{[]string{"127.0.0.1", "147.12.56.11"}, "147.12.56.11"},
	}

	for _, tt := range testData {
		addr, _ := FilterAWS(tt.ips)
		if addr != tt.expected {
			t.Errorf("unexpected result, want: '%s', got: '%s'", tt.expected, addr)
		}
	}
}
