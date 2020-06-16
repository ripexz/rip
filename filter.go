package rip

import (
	"net"
	"strings"
)

var cidrs []*net.IPNet

func init() {
	maxCidrBlocks := []string{
		"127.0.0.1/8",    // localhost
		"10.0.0.0/8",     // 24-bit block
		"172.16.0.0/12",  // 20-bit block
		"192.168.0.0/16", // 16-bit block
		"169.254.0.0/16", // link local address
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	cidrs = make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		cidrs[i] = cidr
	}
}

// Filter allows customisation for detecting or removing specific IPs.
type Filter func(ips []string) (string, bool)

// FilterDefault is used if no Filter function is passed.
func FilterDefault(ips []string) (string, bool) {
	return FilterPublicAddress(ips)
}

// FilterPublicAddress returns the first address that is not under private CIDR blocks.
// List of private CIDR blocks can be seen on:
//
// https://en.wikipedia.org/wiki/Private_network
//
// https://en.wikipedia.org/wiki/Link-local_address
func FilterPublicAddress(ips []string) (string, bool) {
	for _, address := range ips {
		address = strings.TrimSpace(address)
		ipAddress := net.ParseIP(address)
		if ipAddress == nil {
			continue
		}

		var isPrivate bool
		for i := range cidrs {
			if cidrs[i].Contains(ipAddress) {
				isPrivate = true
			}
		}

		if isPrivate {
			continue
		}

		return address, true
	}

	return "", false
}

// FilterAWS returns the last address in the X-Forwarded-For chain.
// This is intended for services behind AWS ELB as it appends the client
// IP address if one is already provided, to prevent spoofing.
// For more info see:
// https://github.com/awsdocs/elb-classic-load-balancers-user-guide/commit/0aa03ef9e048169e18e98d9cc76b32a24df6d29e#diff-539d8b2a6492a6efbfe0ad69b9d61bbf
func FilterAWS(ips []string) (string, bool) {
	if len(ips) == 0 {
		return "", false
	}
	return ips[len(ips)-1], true
}
