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

		if err != nil {
			return ip, fmt.Errorf("Ошибка парсинга IP")
		}
	}

	return ip, nil
}
func GetUserAgent(r *http.Request) (string, error) {
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" {
		return "", fmt.Errorf("Ошибка парсинга User-Agent")
	}

	return userAgent, nil
}
func GetTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Токен не найден")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", fmt.Errorf("Неверный формат токена")
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		return "", fmt.Errorf("Токен не найден")
	}

	return token, nil
}
