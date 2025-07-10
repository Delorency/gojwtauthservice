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
	if exists == 1 {
		return nil, fmt.Errorf("Blocked token")
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

	if payload.UserAgent != data.UserAgent {
		ls := schemes.LogoutRequest{Access: data.Access, Ip: data.Ip, UserAgent: data.UserAgent}
		err := as.Logout(&ls)
		if err != nil {
			return nil, fmt.Errorf("Logount error")
		}
		return nil, fmt.Errorf("Strange UserAgent")
	}

	if payload.Ip != data.Ip {
		go tools.SendToWebhook(data.Ip, as.wburl.WBURL)
	}

	bcrypthash, err := tools.GetBcryptHash(data.Refresh)
	if err != nil {
		return nil, err
	}
	fmt.Println(4)
	obj, err := as.repo.GetByToken(bcrypthash)

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
	access, err := tools.GetJWTToken(as.cfg, obj.UserID, obj.Jti, obj.Ip, obj.UserAgent, obj.User.Email)
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
