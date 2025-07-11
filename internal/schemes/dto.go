package schemes

import (
	"time"

	"github.com/go-playground/validator"
)

func init() {
	_ = validator.New()
}

type AccessCreate struct {
	Jti       string
	Refresh   string
	ExpiredAt time.Time
	Ip        string
	UserAgent string
	UserID    uint
}

type AccessResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	Refresh   string `json:"refresh" validate:"required"`
	Access    string
	Ip        string
	UserAgent string
}

type MeRequest struct {
	Access    string
	Ip        string
	UserAgent string
}

type MeResponse struct {
	Guid uint `json:"guid"`
}

type LogoutRequest struct {
	Access    string
	Ip        string
	UserAgent string
}

type WebhookPayload struct {
	Message string `json:"message"`
}
