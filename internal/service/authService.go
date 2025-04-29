package service

import (
	db "auth/internal/DB/authDB"
	"auth/internal/config"
	"auth/internal/schemes"
)

type AuthServiceI interface {
	Access(*schemes.AccessCreate) (*schemes.AccessResponse, error)
	Refresh(string)
}

type authService struct {
	repo db.AuthDBI
	cfg  *config.ConfigJWTToken
}

func NewAuthService(repo db.AuthDBI, cfg *config.ConfigJWTToken) AuthServiceI {
	return &authService{repo, cfg}
}
