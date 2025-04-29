package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"time"

	"github.com/google/uuid"
)

func (as *authService) Access(data *schemes.AccessCreate) (*schemes.AccessResponse, error) {
	data.Jti = uuid.NewString()
	data.ExpiredAt = time.Now().Add(as.cfg.Rtl)
	data.Refresh = tools.EncodeToBase64()

	refresh, err := as.repo.Access(data)

	if err != nil {
		return nil, err
	}

	return &schemes.AccessResponse{
		AccessToken:  tools.GetJWTToken(),
		RefreshToken: refresh,
	}, nil
}
