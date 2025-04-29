package tools

import (
	"net"
	"net/http"
)

func GetIp(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)

		return ip, err
	}

	return ip, nil
}
