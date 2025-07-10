package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"context"
	"fmt"
	"strings"
	"time"
)

func (as *authService) Logout(data *schemes.LogoutRequest) error {
	key := fmt.Sprintf("used_token:%s", data.Access)

	exists, err := as.redis.Exists(context.Background(), key).Result()
	if exists == 1 {
		return fmt.Errorf("Blocked token")
	}

	// валидный ли access токен?
	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return fmt.Errorf("Invalid token")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])

	if !f {
		return fmt.Errorf("Parse token payload error")
	}

	if payload.UserAgent != data.UserAgent && payload.Ip != data.Ip {
		return fmt.Errorf("Incorrect token")
	}

	obj, f, err := as.repo.GetByUserIDIPUserAgent(payload.Id, payload.Ip, payload.UserAgent)

	if err != nil {
		return fmt.Errorf("Error")
	}

	if !f {
		return fmt.Errorf("Not found")
	}

	// просрочен ли refresh токен
	if time.Now().After(obj.ExpiredAt) {
		return fmt.Errorf("Session expired")
	}

	// jti refresh токена == jti access токена
	if payload.Jti != obj.Jti {
		return fmt.Errorf("Token identificators are not equals")
	}

	// удаляем текущую сессию
	err = as.repo.Delete(obj.ID)
	if err != nil {
		return fmt.Errorf("Delete error")
	}

	err = as.redis.Set(context.Background(), key, "1", as.cfg.Atl).Err()
	if err != nil {
		return fmt.Errorf("Set to redis error")
	}

	return nil
}
