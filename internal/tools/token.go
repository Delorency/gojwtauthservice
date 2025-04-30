package tools

import (
	"auth/internal/config"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Header struct {
	Type string
	Alg  string
}
type Payload struct {
	Iss   string
	Jti   string
	Iat   int64
	Exp   int64
	Ip    string
	Email string
}

func GetJWTToken(cfg *config.ConfigJWTToken, jti string, ip string, email string) (string, error) {
	header := Header{Type: cfg.Typ, Alg: cfg.Alg}
	payload := Payload{Iss: cfg.Iss, Iat: time.Now().Unix(), Exp: time.Now().Add(cfg.Atl).Unix(), Jti: jti, Ip: ip, Email: email}

	header_byte, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	payload_byte, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	base64Header := GetBase64Token(header_byte)
	base64Payload := GetBase64Token(payload_byte)

	data := GetBase64Token(header_byte) + "." + GetBase64Token(payload_byte)

	return fmt.Sprintf("%s.%s.%s", base64Header, base64Payload, GetHmacSha512(data, cfg.SecretKey)), nil
}

func GetHmacSha512(data string, sk string) string {
	h := hmac.New(sha512.New, []byte(sk))
	h.Write([]byte(data))
	signature := h.Sum(nil)

	sign := hex.EncodeToString(signature)

	return sign
}

func GetRefershToken() string {
	return GetBase64Token([]byte(uuid.NewString()))
}
func GetBase64Token(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func GetTokenPayload(block string) (*Payload, bool) {
	pbyte, err := base64.StdEncoding.DecodeString(block)
	if err != nil {
		return nil, false
	}
	p := &Payload{}

	err = json.Unmarshal(pbyte, p)
	if err != nil {
		return nil, false
	}

	return p, true
}

func ValidToken(token string, sk string) bool {
	parts := strings.Split(token, ".")
	hp := strings.Join(parts[0:2], ".")

	if !hmac.Equal([]byte(GetHmacSha512(hp, sk)), []byte(parts[2])) {
		return false
	}

	payload := parts[1]

	p, f := GetTokenPayload(payload)

	if !f {
		return false
	}
	return time.Now().Before(time.Unix(p.Exp, 0))
}
