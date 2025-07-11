package service

import (
	"auth/internal/schemes"
	"auth/internal/tools"
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

func (as *authService) Logout(data *schemes.LogoutRequest) error {
	key := fmt.Sprintf("used_token:%s", data.Access)
	exists, err := as.redis.Exists(context.Background(), key).Result()
	if err != nil {
		log.Println("Check key existence error")
	}
	if exists == 1 {
		return fmt.Errorf("Токен заблокирован")
	}
	// валидный ли access токен?
	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return fmt.Errorf("Токен не валидный")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])

	if !f {
		return fmt.Errorf("Ошибка чтения тела токена")
	}

	if payload.UserAgent != data.UserAgent || payload.Ip != data.Ip {
		return fmt.Errorf("Неверный токен")
	}

	obj, f, err := as.repo.GetByUserIDIPUserAgent(payload.UserID, payload.Ip, payload.UserAgent)

	if err != nil {
		return fmt.Errorf("Ошибка поиска сессии")
	}

	if !f {
		return fmt.Errorf("Сессия не найдена")
	}

	// просрочен ли refresh токен
	if time.Now().After(obj.ExpiredAt) {
		return fmt.Errorf("Refresh токен просрочен")
	}

	// jti refresh токена == jti access токена
	if payload.Jti != obj.Jti {
		return fmt.Errorf("Неверный токен")
	}

	// удаляем текущую сессию
	err = as.repo.Delete(obj.ID)
	if err != nil {
		return fmt.Errorf("Ошибка логаута")
	}

	err = as.redis.Set(context.Background(), key, "1", as.cfg.Atl).Err()
	if err != nil {
		return fmt.Errorf("Set to redis error")
	}

	return nil
}
