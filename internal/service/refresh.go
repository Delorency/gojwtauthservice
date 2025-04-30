package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"fmt"
	"strings"
	"time"
)

func (as *authService) Refresh(data *schemes.RefreshRequest) (*schemes.AccessResponse, error) {

	// валидный ли access токен?
	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return nil, fmt.Errorf("")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])
	if !f {
		return nil, fmt.Errorf("")
	}

	if payload.Ip != data.Ip {
		tools.SendMail(
			payload.Email,
			"Warning",
			"Попытка вход/вход с IP: "+data.Ip,
			as.smtp,
		)
	}

	obj, err := as.repo.AuthorizedUserToken(data.Refresh)

	if err != nil {
		return nil, err
	}

	// просрочен ли refresh токен
	if time.Now().After(obj.ExpiredAt) {
		return nil, fmt.Errorf("")
	}

	// jti refresh токена == jti access токена
	if payload.Jti != obj.Jti {
		return nil, fmt.Errorf("")
	}
	// совпадает ли версия токена с последней версией
	if obj.User.TokenVersion != obj.TokenVersion {
		return nil, fmt.Errorf("")
	}

	// новый access токен
	access, err := tools.GetJWTToken(as.cfg, obj.Jti, obj.Ip, obj.User.Email)
	if err != nil {
		return nil, err
	}

	return &schemes.AccessResponse{
		AccessToken:  access,
		RefreshToken: data.Refresh,
	}, nil
}
