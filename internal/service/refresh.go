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

func (as *authService) Refresh(data *schemes.RefreshRequest) (*schemes.AccessResponse, error) {
	key := fmt.Sprintf("used_token:%s", data.Access)
	exists, err := as.redis.Exists(context.Background(), key).Result()
	if err != nil {
		log.Println("Check key existence error")
	}
	if exists == 1 {
		return nil, fmt.Errorf("Токен заблокирован")
	}

	// валидный ли access токен?
	if !tools.ValidToken(data.Access, as.cfg.SecretKey) {
		return nil, fmt.Errorf("Токен не валидный")
	}

	parts := strings.Split(data.Access, ".")
	payload, f := tools.GetTokenPayload(parts[1])
	if !f {
		return nil, fmt.Errorf("Ошибка чтения тела токена")
	}

	if payload.UserAgent != data.UserAgent {
		ls := schemes.LogoutRequest{Access: data.Access, Ip: data.Ip, UserAgent: data.UserAgent}
		err := as.Logout(&ls)
		if err != nil {
			return nil, fmt.Errorf("Ошибка выхода")
		}
		return nil, fmt.Errorf("Неверное устройство")
	}

	if payload.Ip != data.Ip {
		go tools.SendToWebhook(data.Ip, as.wburl.WBURL)
	}

	obj, f, err := as.repo.GetByUserIDIPUserAgent(payload.UserID, data.Ip, data.UserAgent)
	if err != nil {
		return nil, fmt.Errorf("Ошибка поиска сессии")
	}
	if !f {
		return nil, fmt.Errorf("Сессия не найдена")
	}
	// проверка refresh в теле == refresh в таблице
	if !tools.CheckBcryptHash(data.Refresh, obj.Refresh) {
		return nil, fmt.Errorf("Неверный refresh токен")
	}

	// просрочен ли refresh токен
	if time.Now().After(obj.ExpiredAt) {
		return nil, fmt.Errorf("Refresh токен просрочен")
	}

	// jti refresh токена == jti access токена
	if payload.Jti != obj.Jti {
		return nil, fmt.Errorf("Токены не в паре")
	}

	// новый access токен
	access, err := tools.GetJWTToken(as.cfg, obj.UserID, obj.Jti, obj.Ip, obj.UserAgent)
	if err != nil {
		return nil, err
	}

	err = as.redis.Set(context.Background(), key, "1", as.cfg.Atl).Err()
	if err != nil {
		log.Println("Send to redis error")
	}

	return &schemes.AccessResponse{
		AccessToken:  access,
		RefreshToken: data.Refresh,
	}, nil
}
