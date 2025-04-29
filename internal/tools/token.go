package tools

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func GetJWTToken() (string, error) {
	return "", nil
}

func EncodeToBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
}
