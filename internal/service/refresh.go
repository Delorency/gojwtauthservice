package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"fmt"
	"strings"
	"time"
)

func (as *authService) Refresh(data *schemes.RefreshRequest) (*schemes.AccessResponse, error) {
	obj, f, err := as.repo.AuthorizedUserAgentToken(data.Refresh, data.UserAgent)

	if err != nil {
		return nil, err
	}

	// такой refresh под таким устройством на найден
	if !f {
		return nil, fmt.Errorf("")
	}

	// refresh просрочен?
	if !time.Now().Before((obj.ExpiredAt)) {
		return nil, fmt.Errorf("")
	}

	// совпадает ли версия токена с последней версией
	if obj.User.TokenVersion != obj.TokenVersion {
		return nil, fmt.Errorf("")
	}

	// валидный ли access токен?
	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return nil, fmt.Errorf("")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])
	if !f {
		return nil, fmt.Errorf("")
	}

	// jti refresh токена == jti access токена
	if payload.Jti != obj.Jti {
		return nil, fmt.Errorf("")
	}

	// новый access токен
	access, err := tools.GetJWTToken(as.cfg, obj.Jti)
	if err != nil {
		return nil, err
	}

	return &schemes.AccessResponse{
		AccessToken:  access,
		RefreshToken: data.Refresh,
	}, nil
} // ZGRkMWQ1MTctYmMzNS00ZDQ3LWFjYzMtYWFkZmYwNWZiNzhi
