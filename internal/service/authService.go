package service

import (
	db "auth/internal/DB/authDB"
	userdb "auth/internal/DB/userDB"
	"auth/internal/config"
	"auth/internal/schemes"

	"github.com/redis/go-redis/v9"
)

type AuthServiceI interface {
	Access(*schemes.AccessCreate) (*schemes.AccessResponse, error)
	Refresh(*schemes.RefreshRequest) (*schemes.AccessResponse, error)
	Logout(*schemes.LogoutRequest) error
	Me(data *schemes.MeRequest) (uint, error)
}

type authService struct {
	repo   db.AuthDBI
	redis  *redis.Client
	userdb userdb.UserDBI
	cfg    *config.ConfigJWTToken
	wburl  *config.ConfigWebhook
}

func NewAuthService(repo db.AuthDBI, redis *redis.Client, userdb userdb.UserDBI, cfg *config.ConfigJWTToken, webhook *config.ConfigWebhook) AuthServiceI {
	return &authService{repo, redis, userdb, cfg, webhook}
}
