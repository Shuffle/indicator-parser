package ioc

import (
	"net"
	"strings"
)

// for both IPv4 and IPv6
func IsIp(value string) bool {
	return net.ParseIP(value) != nil
}

func IsUrl(value string) bool {
	return strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://")
}

func IsEmail(value string) bool {
	vals := strings.Split(value, "@")

	switch {
	case len(vals) != 2:
		return false
	case vals[0] == "":
		return false
	case IsDomain(vals[1]):
		return true
	}
	return false
}

// refactor?
func IsDomain(value string) bool {
	if IsIp(value) {
		return false
	}

	if IsUrl(value) {
		return true
	}

	if IsEmail(value) {
		return false
	}

	url, err := ParseURL("http://" + value)
	if err != nil {
		return false
	}

	if len(strings.Split(url.Host, ".")) >= 2 {
		return true
	}

	return false
}
