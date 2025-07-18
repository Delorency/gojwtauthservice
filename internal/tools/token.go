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
	"golang.org/x/crypto/bcrypt"
)

type Header struct {
	Type string
	Alg  string
}
type Payload struct {
	Iss       string
	Jti       string
	Iat       int64
	Exp       int64
	UserID    uint
	Ip        string
	UserAgent string
}

func GetJWTToken(cfg *config.ConfigJWTToken, id uint, jti, ip, useragent string) (string, error) {
	header := Header{Type: cfg.Typ, Alg: cfg.Alg}
	payload := Payload{
		Iss:       cfg.Iss,
		Iat:       time.Now().Unix(),
		Exp:       time.Now().Add(cfg.Atl).Unix(),
		Jti:       jti,
		UserID:    id,
		Ip:        ip,
		UserAgent: useragent,
	}

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

	data := base64Header + "." + base64Payload

	return fmt.Sprintf("%s.%s", data, GetHmacSha512(data, cfg.SecretKey)), nil
}

func GetHmacSha512(data string, sk string) string {
	h := hmac.New(sha512.New, []byte(sk))
	h.Write([]byte(data))
	signature := h.Sum(nil)

	sign := hex.EncodeToString(signature)

	return sign
}

func GetBase64Token(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func GetRefreshToken() string {
	return GetBase64Token([]byte(uuid.NewString()))
}

func GetBcryptHash(data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
func CheckBcryptHash(refreshToken, bcryptHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptHash), []byte(refreshToken))
	return err == nil
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
