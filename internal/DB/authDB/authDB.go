package authdb

import (
	userdb "auth/internal/DB/userDB"
	"auth/internal/models"
	"auth/internal/schemes"

	"gorm.io/gorm"
)

type AuthDBI interface {
	AuthorizedUserAgent(uint, string) (*models.RefreshToken, bool, error)
	AuthorizedUserToken(string) (*models.RefreshToken, error)
	Create(*schemes.AccessCreate) error
	Update(*models.RefreshToken) error
}

type authDB struct {
	db     *gorm.DB
	userdb userdb.UserDBI
}

func NewAuthDB(db *gorm.DB, userdb userdb.UserDBI) AuthDBI {
	return &authDB{db, userdb}
}
