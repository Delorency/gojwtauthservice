package authdb

import (
	userdb "auth/internal/DB/userDB"
	"auth/internal/models"
	"auth/internal/schemes"

	"gorm.io/gorm"
)

type AuthDBI interface {
	AuthorizedUserAgent(uint, string) (*models.RefreshToken, bool, error)
	AuthorizedUserAgentToken(string, string) (*models.RefreshToken, bool, error)
	Create(*schemes.AccessCreate) (string, error)
	Update(*models.RefreshToken) error
}

type authDB struct {
	db     *gorm.DB
	userdb userdb.UserDBI
}

func NewAuthDB(db *gorm.DB, userdb userdb.UserDBI) AuthDBI {
	return &authDB{db, userdb}
}
