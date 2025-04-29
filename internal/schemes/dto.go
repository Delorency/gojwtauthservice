package schemes

import "time"

type AccessCreate struct {
	Jti          string
	UserAgent    string
	Refresh      string
	ExpiredAt    time.Time
	Ip           string
	UserID       uint
	TokenVersion string
}

type AccessResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	Refresh string `json:"refresh"`
	Ip      string `json:"ip"`
}
