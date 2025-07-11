package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"context"
	"fmt"
	"log"
	"strings"
)

func (as *authService) Me(data *schemes.MeRequest) (uint, error) {
	key := fmt.Sprintf("used_token:%s", data.Access)
	exists, err := as.redis.Exists(context.Background(), key).Result()
	if err != nil {
		log.Println("Check key existence error")
	}
	if exists == 1 {
		return 0, fmt.Errorf("Токен заблокирован")
	}

	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return 0, fmt.Errorf("Токен не валидный")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])

	if !f {
		return 0, fmt.Errorf("Ошибка чтения тела токена")
	}

	if payload.UserAgent != data.UserAgent || payload.Ip != data.Ip {
		return 0, fmt.Errorf("Неверный токен")
	}

	return payload.UserID, nil
}
