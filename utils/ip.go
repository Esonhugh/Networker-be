package utils

import (
	"net"
	"regexp"
)

func IsGoodIP(addr string) bool {
	success := net.ParseIP(addr)
	if success == nil {
		return false
	}
	if success.IsLoopback() {
		return false
	}
	if success.IsMulticast() {
		return false
	}
	return true
}

const fqdn = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$`

func IsGoodDoamin(domain string) bool {
	return regexp.MustCompile(fqdn).MatchString(domain)
}
