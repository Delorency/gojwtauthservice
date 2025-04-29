package tools

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func GetJWTToken() (string, error) {
	return "", nil
}

func EncodeToBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
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
