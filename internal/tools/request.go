package tools

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func GetIp(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)

		return ip, err
	}

	return ip, nil
}
func GetAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header missing")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", fmt.Errorf("Invalid token format")
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		return "", fmt.Errorf("Token is empty")
	}

	return token, nil
}
