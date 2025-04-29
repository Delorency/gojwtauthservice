package service

import (
	db "auth/internal/DB/authDB"
	userdb "auth/internal/DB/userDB"
	"auth/internal/config"
	"auth/internal/schemes"
)

type AuthServiceI interface {
	Access(*schemes.AccessCreate) (*schemes.AccessResponse, error)
	Refresh(string)
}

type authService struct {
	repo   db.AuthDBI
	userdb userdb.UserDBI
	cfg    *config.ConfigJWTToken
}

func NewAuthService(repo db.AuthDBI, userdb userdb.UserDBI, cfg *config.ConfigJWTToken) AuthServiceI {
	return &authService{repo, userdb, cfg}
}
