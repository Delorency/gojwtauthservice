package tools

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func GetJWTToken() string {
	return ""
}

func EncodeToBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))
}
