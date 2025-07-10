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
}

type authService struct {
	repo   db.AuthDBI
	redis  *redis.Client
	userdb userdb.UserDBI
	cfg    *config.ConfigJWTToken
	smtp   *config.ConfigSMTP
}

func NewAuthService(repo db.AuthDBI, redis *redis.Client, userdb userdb.UserDBI, cfg *config.ConfigJWTToken, smtp *config.ConfigSMTP) AuthServiceI {
	return &authService{repo, redis, userdb, cfg, smtp}
}
